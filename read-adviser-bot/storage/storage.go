package storage

import (
	"crypto/sha256"
	"e"
	"errors"
	"fmt"
)

var (
	ErrNoSavedLinks = errors.New("no saved pages")
)

type IStorage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha256.New()

	if _, err := h.Write([]byte(p.URL + p.UserName)); err != nil {
		return "", e.WrapIfErr("failed to calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
