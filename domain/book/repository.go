package book

import "context"

// Repository represent the User's repository contract.
type Repository interface {
	Create(ctx context.Context, book *Book) error
	Get(ctx context.Context, bookID string) (*Book, error)
	GetAll(ctx context.Context, book *Book) ([]Book, error)
}
