package postgresql

import (
	"context"
	"github.com/deeprave/go-auth/lib"
	"github.com/deeprave/go-auth/models"
	"github.com/deeprave/go-auth/repository"
)

func (pg *PG) GetUsers(activeOnly bool, window ...*repository.Window) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(pg.CTX, pgTimeout)
	defer cancel()

	var win *repository.Window = nil
	args := lib.NewArgs()
	if activeOnly {
		args = args.Append("WHERE is_active")
	}
	if len(window) > 0 {
		win = window[0]
		args = args.Append(win.ToString())
		win.Next()
	}

	query := models.UserTable.Select(args)

	rows, err := pg.DB.QueryContext(ctx, query)
	if err == nil {
		defer Close(rows)

		var users []*models.User
		for rows.Next() {
			var user models.User
			if err = rows.Scan(user.ScanFields()...); err != nil {
				break
			}
			users = append(users, &user)
		}
		return users, nil
	}
	return nil, err
}

func (pg *PG) GetUserById(id int64) (*models.User, error) {
	ctx, cancel := context.WithTimeout(pg.CTX, pgTimeout)
	defer cancel()

	args := lib.NewArgs("WHERE id = $1")

	query := models.UserTable.Select(args)
	row := pg.DB.QueryRowContext(ctx, query, id)
	var (
		user models.User
		err  error
	)
	if err = row.Scan(user.ScanFields()...); err == nil {
		return &user, nil
	}
	return nil, err
}

func (pg *PG) GetUserByName(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(pg.CTX, pgTimeout)
	defer cancel()

	args := lib.NewArgs("WHERE username = $1")

	query := models.UserTable.Select(args)
	row := pg.DB.QueryRowContext(ctx, query, username)
	var (
		user models.User
		err  error
	)
	if err = row.Scan(user.ScanFields()...); err == nil {
		return &user, nil
	}
	return nil, err
}

func (pg *PG) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(pg.CTX, pgTimeout)
	defer cancel()

	args := lib.NewArgs("WHERE email = $1")

	query := models.UserTable.Select(args)
	row := pg.DB.QueryRowContext(ctx, query, email)
	var (
		user models.User
		err  error
	)
	if err = row.Scan(user.ScanFields()...); err == nil {
		return &user, nil
	}
	return nil, err
}
