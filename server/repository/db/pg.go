package db

import (
	"context"
	"database/sql"
	"github.com/deeprave/go-auth/models"
	"github.com/deeprave/go-auth/repository"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"io"
	"log"
	"time"
)

type PG struct {
	DB  *sql.DB
	CTX context.Context
}

const pgTimeout = time.Second * 5

func Close(obj io.Closer) {
	_ = obj.Close()
}

func NewPG(dsn string) (*PG, error) {
	var (
		pg  = PG{}
		err error
	)

	if pg.DB, err = sql.Open("pgx", dsn); err == nil {
		if err = pg.DB.Ping(); err == nil {
			log.Println("successfully connected to database")
			pg.CTX = context.Background()
		} else {
			log.Println("error connecting to database:", err)
		}
	}
	return &pg, err
}

func (pg *PG) Close() {
	err := pg.DB.Close()
	if err != nil {
		log.Println("error closing connection to database:", err)
	}
}

func (pg *PG) GetUsers(activeOnly bool, window ...*repository.Window) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(pg.CTX, pgTimeout)
	defer cancel()

	var win *repository.Window = nil
	args := repository.NewArgs()
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

	args := repository.NewArgs("WHERE id = $1")

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

	args := repository.NewArgs("WHERE username = $1")

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

	args := repository.NewArgs("WHERE email = $1")

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

func (pg *PG) GetCredentialsForUser(id int64) ([]*models.Credential, error) {
	ctx, cancel := context.WithTimeout(pg.CTX, pgTimeout)
	defer cancel()

	args := repository.NewArgs("WHERE user_id = $1")

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
