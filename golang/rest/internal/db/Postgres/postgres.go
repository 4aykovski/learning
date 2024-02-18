package Postgres

import (
	"database/sql"
	"fmt"

	"github.com/4aykovski/learning/golang/rest/internal/config"
	"github.com/4aykovski/learning/golang/rest/internal/db"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func New(cfg config.Postgres) (*Postgres, error) {
	pq, err := sql.Open("postgres", cfg.DSNTemplate)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", db.ErrCantCreateDatabase, err)
	}

	if err = pq.Ping(); err != nil {
		return nil, fmt.Errorf("%w: %w", db.ErrCantPingDatabase, err)
	}

	if err = databasePrepare(pq); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &Postgres{db: pq}, nil
}

func databasePrepare(pq *sql.DB) error {
	stmt1, err := pq.Prepare(`
	CREATE TABLE IF NOT EXISTS url(
	    id INTEGER PRIMARY KEY,
	    alias TEXT NOT NULL UNIQUE,
	    url TEXT NOT NULL
	);`)
	if err != nil {
		return fmt.Errorf("%w: %w", db.ErrCantPrepareDatabase, err)
	}

	_, err = stmt1.Exec()
	if err != nil {
		return fmt.Errorf("%w: %w", db.ErrCantPrepareDatabase, err)
	}

	stmt2, err := pq.Prepare(`
	CREATE INDEX idx_alias ON url(alias);`)

	if err != nil {
		return fmt.Errorf("%w: %w", db.ErrCantPrepareDatabase, err)
	}

	_, err = stmt2.Exec()
	if err != nil {
		return fmt.Errorf("%w: %w", db.ErrCantPrepareDatabase, err)
	}

	return nil
}
