package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	domain "github.com/wackGarcia/books/domain/user"
)

type Repository struct {
}

func NewUserMockRepository() *Repository {
	return &Repository{}
}

func (store *Repository) Create(ctx context.Context, user *domain.User) error {
	return errors.New("some error")
}

func (store *Repository) Get(ctx context.Context, username string) (*domain.User, error) {

	return &domain.User{
		ID:        uuid.NewString(),
		UserName:  username,
		Password:  "123456",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}, nil
}
