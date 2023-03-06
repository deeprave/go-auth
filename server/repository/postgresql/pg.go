package postgresql

import (
	"context"
	"database/sql"
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
