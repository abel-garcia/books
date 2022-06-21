package user

import "context"

// UserService represent the User's usecases.
type UserService interface {
	Create(ctx context.Context, user *User) error
	Get(ctx context.Context, username string) (*User, error)
}

type Service struct {
	repository UserService
}

func NewUserService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) Create(ctx context.Context, user *User) error {
	return service.repository.Create(ctx, user)
}

func (service *Service) Get(ctx context.Context, username string) (*User, error) {
	return service.repository.Get(ctx, username)
}
