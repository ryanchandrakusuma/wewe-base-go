package repository

import (
	"context"
	"ryanchandrakusuma/wewe-base-go/domain"

	"github.com/jackc/pgx/v5"
)

type bookRepository struct {
	Conn *pgx.Conn
}

func NewBookRepository(conn *pgx.Conn) domain.BookRepository {
	return &bookRepository{conn}
}

func (r *bookRepository) Fetch(ctx context.Context) (result []domain.Book, err error) {
	query := "SELECT id, title, author, quantity FROM books"
    rows, err := r.Conn.Query(context.Background(), query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var book domain.Book
        if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Quantity); err != nil {
            return nil, err
        }
        result = append(result, book)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return result, nil
}


