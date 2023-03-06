package postgresql

import (
	"context"
	"github.com/deeprave/go-auth/models"
)

func (pg *PG) GetDatabaseInfo() (*models.Database, error) {
	ctx, cancel := context.WithTimeout(pg.CTX, pgTimeout)
	defer cancel()

	info, err := pg.getDatabase(ctx)
	if err == nil {
		// collect schemas
		info.Schemas, err = pg.getSchemas(ctx)
		if err == nil {
			// collect tables for each schema
			for _, schema := range info.Schemas {
				schema.Tables, err = pg.getTablesForSchema(ctx, schema)
				if err == nil {
					for _, table := range schema.Tables {
						table.Fields, err = pg.getFieldsForTable(ctx, table)
						if err == nil {
							err = pg.getTableData(ctx, &table)
						}
						if err != nil {
							break
						}
					}
				}
				if err != nil {
					break
				}
			}
		}
	}
	return info, err
}

func (pg *PG) getDatabase(ctx context.Context) (*models.Database, error) {
	info := models.Database{}

	// database name
	row := pg.DB.QueryRowContext(ctx, "SELECT current_database()")
	err := row.Scan(&info.Database)
	if err == nil {
		// general database stats
		row = pg.DB.QueryRowContext(ctx, "SELECT xact_commit, xact_rollback, sessions, deadlocks from pg_stat_database WHERE datname = $1", info.Database)
		err = row.Scan(&info.Commits, &info.Rollbacks, &info.Sessions, &info.Deadlocks)
	}
	return &info, err
}

func (pg *PG) getSchemas(ctx context.Context) ([]models.Schema, error) {
	rows, err := pg.DB.QueryContext(ctx, "SELECT oid, nspname from pg_namespace WHERE nspname <> 'information_schema' AND NOT nspname LIKE 'pg_%'")
	var schemas []models.Schema
	if err == nil {
		defer Close(rows)
		for rows.Next() {
			schema := models.Schema{}
			err = rows.Scan(&schema.Oid, &schema.Name)
			if err != nil {
				break
			}
			schemas = append(schemas, schema)
		}
	}
	return schemas, err
}

func (pg *PG) getTablesForSchema(ctx context.Context, schema models.Schema) ([]models.Table, error) {
	rows, err := pg.DB.QueryContext(ctx, "SELECT oid, relname from pg_class WHERE relkind = 'r' AND relnamespace = $1", schema.Oid)
	var tables []models.Table
	if err == nil {
		defer Close(rows)
		for rows.Next() {
			table := models.Table{}
			err = rows.Scan(&table.Oid, &table.Name)
			if err != nil {
				break
			}
			tables = append(tables, table)
		}
	}
	return tables, err
}

func (pg *PG) getFieldsForTable(ctx context.Context, table models.Table) ([]models.Field, error) {
	rows, err := pg.DB.QueryContext(ctx, "SELECT attname, atttypid::regtype, "+
		"CASE WHEN attlen > 0 THEN attlen ELSE CASE WHEN atttypmod > 0 THEN atttypmod ELSE -1 END END "+
		"FROM pg_attribute "+
		"WHERE attrelid = $1 AND attnum > 0 AND NOT attisdropped "+
		"ORDER BY attnum", table.Oid)
	var fields []models.Field
	if err == nil {
		defer Close(rows)
		for rows.Next() {
			field := models.Field{}
			err = rows.Scan(&field.Name, &field.Type, &field.Length)
			if err != nil {
				break
			}
			fields = append(fields, field)
		}
	}
	return fields, err
}

func (pg *PG) getTableData(ctx context.Context, table *models.Table) error {
	row := pg.DB.QueryRowContext(ctx, "SELECT count(*) FROM "+table.TableName())
	err := row.Scan(&table.NumRecords)
	if err == nil {
		table.NumAvailable = table.NumRecords
		if table.HasFieldName("dt_deleted") {
			row = pg.DB.QueryRowContext(ctx, "SELECT count(*) FROM "+table.TableName()+" WHERE dt_deleted IS NOT NULL")
			err = row.Scan(&table.NumAvailable)
		}
		if err == nil {
			if table.NumRecords > 0 {
				if table.HasFieldName("id") {
					row = pg.DB.QueryRowContext(ctx, "SELECT id FROM "+table.TableName()+" ORDER BY id ASC LIMIT 1")
					err = row.Scan(&table.FirstId)
					if err == nil {
						row = pg.DB.QueryRowContext(ctx, "SELECT id FROM "+table.TableName()+" ORDER BY id DESC LIMIT 1")
						err = row.Scan(&table.LastId)
					}
				}
				if err == nil && table.HasFieldName("dt_created") {
					row = pg.DB.QueryRowContext(ctx, "SELECT dt_created FROM "+table.TableName()+" ORDER BY dt_created ASC LIMIT 1")
					err = row.Scan(&table.FirstCreated)
					if err == nil {
						row = pg.DB.QueryRowContext(ctx, "SELECT dt_created FROM "+table.TableName()+" ORDER BY dt_created DESC LIMIT 1")
						err = row.Scan(&table.LastCreated)
					}
				}
				if err == nil && table.HasFieldName("dt_updated") {
					row = pg.DB.QueryRowContext(ctx, "SELECT dt_updated FROM "+table.TableName()+" ORDER BY dt_updated DESC LIMIT 1")
					err = row.Scan(&table.LastUpdated)
				}
			}
		}
	}
	return err
	//rows, err := pg.DB.QueryContext(ctx, "SELECT oid, relname from pg_class WHERE relkind = 'r' AND relnamespace = $1", table.Oid)
	//if err == nil {
	//	defer Close(rows)
	//	for rows.Next() {
	//		err := models.TableDef{}
	//		err = rows.Scan(&table.Name)
	//		if err != nil {
	//			break
	//		}
	//	}
	//}
	//return err
}
