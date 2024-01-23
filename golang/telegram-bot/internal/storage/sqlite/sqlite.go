package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/internal/storage"
	errorWrapper "github.com/4aykovski/learning/tree/main/golang/telegram-bot/lib/error-wrapper"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, errorWrapper.Wrap("can't open database", err)
	}

	if err = db.Ping(); err != nil {
		return nil, errorWrapper.Wrap("can't connect to database", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Save(ctx context.Context, p *storage.Page) error {
	q := `INSERT INTO pages (url, user_name) VALUES (?, ?)`

	_, err := s.db.ExecContext(ctx, q, p.URL, p.UserName)
	if err != nil {
		return errorWrapper.Wrap("can't save a page", err)
	}

	return nil
}

func (s *Storage) PickRandom(ctx context.Context, userName string) (*storage.Page, error) {
	q := `SELECT url FROM pages WHERE user_name = ? ORDER BY RANDOM() LIMIT 1`

	var url string

	err := s.db.QueryRowContext(ctx, q, userName).Scan(&url)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, storage.ErrNoSavedPages
	}
	if err != nil {
		return nil, errorWrapper.Wrap("can't pick a random page", err)
	}

	return &storage.Page{
		URL:      url,
		UserName: userName,
	}, nil
}

func (s *Storage) Remove(ctx context.Context, page *storage.Page) error {
	q := `DELETE FROM pages WHERE url = ? AND user_name = ?`

	_, err := s.db.ExecContext(ctx, q, page.URL, page.UserName)
	if err != nil {
		return errorWrapper.Wrap("can't remove a page", err)
	}

	return nil
}

func (s *Storage) IsExists(ctx context.Context, page *storage.Page) (bool, error) {
	q := `SELECT COUNT(*) FROM pages WHERE url = ? AND user_name = ?`

	var count int

	err := s.db.QueryRowContext(ctx, q, page.URL, page.UserName).Scan(&count)
	if err != nil {
		return false, errorWrapper.Wrap("can't check if page exists", err)
	}

	return count > 0, nil
}

func (s *Storage) Init(ctx context.Context) error {
	q := `CREATE TABLE IF NOT EXISTS pages (url TEXT, user_name TEXT)`

	_, err := s.db.ExecContext(ctx, q)
	if err != nil {
		return errorWrapper.Wrap("can't create a table", err)
	}

	return nil
}
