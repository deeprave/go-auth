package postgresql

import (
	"context"
	"github.com/deeprave/go-testutils/test"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	testSetup()
	code := m.Run()
	testShutdown()
	os.Exit(code)
}

func testSetup() {
}

func testShutdown() {
}

func TestPG_getDatabaseInfo(t *testing.T) {
	url, set := os.LookupEnv("DATABASE_URL")
	if !set {
		t.Fatalf("envvar DATABASE_URL is required for this test")
	}
	pg, err := NewPG(url)
	test.ShouldBeNoError(t, err, "unexpected database error: %v", err)

	ctx, cancel := context.WithCancel(pg.CTX)
	defer cancel()

	// Database metadata
	info, err := pg.getDatabase(ctx)
	test.ShouldBeNoError(t, err, "unexpected database error: %v", err)
	test.ShouldBeEqual(t, info.Database, "auth")
	test.ShouldBeTrue(t, info.Commits > 0, "expected commits > 0")
	test.ShouldBeTrue(t, info.Sessions > 0, "expected sessions > 0")

	// Schemas present
	info.Schemas, err = pg.getSchemas(ctx)
	test.ShouldBeNoError(t, err, "unexpected database error: %v", err)
	test.ShouldBeTrue(t, len(info.Schemas) == 1, "only 1 schema present")
	test.ShouldBeEqual(t, info.Schemas[0].Name, "public")

	for _, schema := range info.Schemas {
		schema.Tables, err = pg.getTablesForSchema(ctx, schema)
		test.ShouldBeNoError(t, err, "unexpected database error: %v", err)
		test.ShouldBeTrue(t, len(schema.Tables) >= 2, "contains at least 2 tables")

		for _, table := range schema.Tables {
			table.Fields, err = pg.getFieldsForTable(ctx, table)
			test.ShouldBeNoError(t, err, "unexpected database error: %v", err)
			test.ShouldBeTrue(t, len(table.Fields) >= 2, "contains at least 3 fields")

			err = pg.getTableData(ctx, &table)
			test.ShouldBeNoError(t, err, "unexpected database error: %v", err)
		}
	}
}
