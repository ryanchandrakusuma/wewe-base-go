package usecase

import (
	"context"
	"ryanchandrakusuma/wewe-base-go/domain"
	"time"
)

type bookUsecase struct {
	bookRepo domain.BookRepository
	contextTimeout time.Duration
}

func NewBookUsecase(b domain.BookRepository, timeout time.Duration) domain.BookUsecase {
	return &bookUsecase{
		bookRepo: b,
		contextTimeout: timeout,
	}
}

// Fetch implements domain.BookUsecase.
func (bu *bookUsecase) Fetch(ctx context.Context) (result []domain.Book, err error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	result, err = bu.bookRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	
	return result, err
}
