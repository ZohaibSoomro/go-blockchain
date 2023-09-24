package models

import (
	"crypto/md5"
	"fmt"
)

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	PublishDate string `json:"publish_date"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
}

func NewBook(t, a, pd, isbn string) *Book {
	return &Book{
		Title: t, Author: a, PublishDate: pd, ISBN: isbn,
	}
}

func (b *Book) GenerateID() string {
	data := []byte(b.Title + b.Author + b.PublishDate + b.ISBN)
	n := md5.New()
	n.Write([]byte(data))
	b.ID = fmt.Sprintf("%x", n.Sum(nil))
	return b.ID
}
