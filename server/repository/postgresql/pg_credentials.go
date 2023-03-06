package postgresql

import (
	"context"
	"github.com/deeprave/go-auth/lib"
	"github.com/deeprave/go-auth/models"
)

func (pg *PG) GetCredentialsForUser(id int64) ([]*models.Credential, error) {
	ctx, cancel := context.WithTimeout(pg.CTX, pgTimeout)
	defer cancel()

	args := lib.NewArgs("WHERE user_id = $1")

	query := models.CredentialTable.Select(args)

	rows, err := pg.DB.QueryContext(ctx, query, id)
	if err == nil {
		defer Close(rows)

		var creds []*models.Credential
		for rows.Next() {
			var cred models.Credential
			if err = rows.Scan(cred.ScanFields()...); err != nil {
				break
			}
			creds = append(creds, &cred)
		}
		return creds, nil
	}
	return nil, err
}
