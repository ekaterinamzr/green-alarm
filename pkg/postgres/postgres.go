package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sqlx.DB
}

const (
	connAttempts = 5
	connTimeout  = 5 * time.Second
)

func New(dsn string) (*Postgres, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Postgres could not open: %w", err)
	}

	for i := 0; i < connAttempts; i++ {
		err = db.Ping()

		if err == nil {
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", connAttempts-i-1)
		time.Sleep(connTimeout)
	}

	if err != nil {
		return nil, fmt.Errorf("Postgres could not connect: %w", err)
	}

	return &Postgres{DB: db}, nil
}

func (pg *Postgres) Close() error {
	return pg.DB.Close()
}
