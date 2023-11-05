package files

import (
	"e"
	"encoding/gob"
	"errors"
	"math/rand"
	"os"
	"path/filepath"
	"storage"
	"time"
)

type Storage struct {
	basePath string
}

// func new storage
func New(basePath string) *Storage {
	return &Storage{
		basePath: basePath,
	}
}

const (
	defaultPerm = 0774
)

var (
	ErrNoSavedPages = errors.New("no saved pages")
)

// func save page to storage
func (s *Storage) Save(page *storage.Page) (err error) {
	defer func() {
		err = e.WrapIfErr("failed to save page", err)
	}()

	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}
	fName, err := fileName(page)

	fPath = filepath.Join(fPath, fName)

	f, err := os.Create(fPath)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	if err := gob.NewEncoder(f).Encode(page); err != nil {
		return err
	}

	return nil
}
func (s *Storage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() {
		err = e.WrapIfErr("failed to pick random page", err)
	}()

	fPath := filepath.Join(s.basePath, userName)
	files, err := os.ReadDir(fPath)

	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, ErrNoSavedPages
	}

	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := randGenerator.Intn(len(files))
	randFile := files[n]

	return s.decodePath(filepath.Join(fPath, randFile.Name()))

}

func (s *Storage) decodePath(fPath string) (page *storage.Page, err error) {
	file, err := os.Open(fPath)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = file.Close()
	}()

	if err := gob.NewDecoder(file).Decode(&page); err != nil {
		return nil, err
	}

	return page, nil
}

func (s *Storage) Remove(page *storage.Page) (err error) {
	defer func() {
		err = e.WrapIfErr("failed to remove file", err)
	}()

	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath := filepath.Join(s.basePath, page.UserName, fName)

	if err := os.Remove(fPath); err != nil {
		return err
	}
	return nil
}

func (s *Storage) IsExists(page *storage.Page) (bool, error) {

	fName, err := fileName(page)
	if err != nil {
		return false, e.Wrap("failed to check file existence", err)
	}
	
	fPath := filepath.Join(s.basePath, page.UserName, fName)

	switch _, err = os.Stat(fPath); {
		case errors.Is(err, os.ErrNotExist):
			return false, nil
		case err != nil:
			return false, e.Wrap("failed to check file existence", err)
	}
	return true, nil
}

func fileName(page *storage.Page) (string, error) {
	return page.Hash()
}
