package book

import "context"

// Repository represent the User's repository contract.
type Repository interface {
	Create(ctx context.Context, book *Book) (string, error)
	GetAll(ctx context.Context, book *Book) ([]Book, error)
	Delete(ctx context.Context, bookID, UserID string) error
}
