package domain

import (
	"context"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

type BookUsecase interface {
	Fetch(ctx context.Context) ([]Book, error)
}

type BookRepository interface {
	Fetch(ctx context.Context) (res []Book, err error)
}

