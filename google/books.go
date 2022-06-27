package google

import (
	"context"

	"google.golang.org/api/books/v1"
)

type Service struct {
	service *books.Service
}

func New() (*Service, error) {
	ctx := context.Background()
	booksService, err := books.NewService(ctx)
	if err != nil {
		return nil, err
	}

	return &Service{
		service: booksService,
	}, nil
}
