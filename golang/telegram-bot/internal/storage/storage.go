package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

	error_wrapper "github.com/4aykovski/learning/tree/main/golang/telegram-bot/lib/error-wrapper"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

var ErrNoSavedPages = errors.New("no saved pages")

type Page struct {
	URL      string
	UserName string
	//Created  time.Time
}

func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", error_wrapper.Wrap("can't generate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", error_wrapper.Wrap("can't generate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}