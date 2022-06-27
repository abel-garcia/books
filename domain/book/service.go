package book

import "context"

// BookService represent the User's usecases.
type BookService interface {
	Create(ctx context.Context, book *Book) (string, error)
	GetAll(ctx context.Context, book *Book) ([]Book, error)
	Delete(ctx context.Context, bookID, UserID string) error
}

type Service struct {
	repository BookService
}

func NewBookService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) Create(ctx context.Context, book *Book) (string, error) {
	return service.repository.Create(ctx, book)
}

func (service *Service) GetAll(ctx context.Context, book *Book) ([]Book, error) {
	return service.repository.GetAll(ctx, book)
}

func (service *Service) Delete(ctx context.Context, bookID, userID string) error {
	return service.repository.Delete(ctx, bookID, userID)
}
