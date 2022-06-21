package book

import "context"

// BookService represent the User's usecases.
type BookService interface {
	Create(ctx context.Context, book *Book) error
	Get(ctx context.Context, bookID string) (*Book, error)
	GetAll(ctx context.Context, book *Book) ([]Book, error)
}

type Service struct {
	repository BookService
}

func NewBookService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) Create(ctx context.Context, book *Book) error {
	return service.repository.Create(ctx, book)
}

func (service *Service) Get(ctx context.Context, bookID string) (*Book, error) {
	return service.repository.Get(ctx, bookID)
}

func (service *Service) GetAll(ctx context.Context, book *Book) ([]Book, error) {
	return service.repository.GetAll(ctx, book)
}
