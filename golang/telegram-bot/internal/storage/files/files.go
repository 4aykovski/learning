package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/internal/storage"
	error_wrapper "github.com/4aykovski/learning/tree/main/golang/telegram-bot/lib/error-wrapper"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = error_wrapper.WrapIfErr("can't save page", err) }()

	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}

	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}

func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() { err = error_wrapper.WrapIfErr("can't pick random page", err) }()

	path := filepath.Join(s.basePath, userName)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return nil, err
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, storage.ErrNoSavedPages
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := rand.Intn(len(files))

	f := files[n]

	return s.decodePage(filepath.Join(path, f.Name()))
}

func (s Storage) Remove(p *storage.Page) error {
	fName, err := fileName(p)
	if err != nil {
		return error_wrapper.Wrap("can't remove a file", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fName)

	if err := os.Remove(path); err != nil {
		msg := fmt.Sprintf("can't remove a file %s", path)
		return error_wrapper.Wrap(msg, err)
	}

	return nil
}

func (s Storage) IsExists(p *storage.Page) (bool, error) {
	fName, err := fileName(p)
	if err != nil {
		return false, error_wrapper.Wrap("can't check if file %s exists", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fName)

	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		msg := fmt.Sprintf("can't check if file %s exists", path)
		return false, error_wrapper.Wrap(msg, err)
	}

	return true, nil
}

func (s Storage) decodePage(fPath string) (*storage.Page, error) {
	f, err := os.Open(fPath)
	if err != nil {
		return nil, error_wrapper.Wrap("can't decode a page", err)
	}

	defer func() { _ = f.Close() }()

	var p storage.Page

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, error_wrapper.Wrap("can't decode a page", err)
	}

	return &p, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
