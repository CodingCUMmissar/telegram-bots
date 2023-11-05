package storage

import (
	"crypto/sha1"
	"e"
	"fmt"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) bool
}

type Page struct {
	URL      string
	UserName string
}


func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := h.Write([]byte(p.URL + p.UserName)); err != nil {
		return "", e.WrapIfErr("failed to calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}