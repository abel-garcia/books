package user

import "context"

// Repository represent the User's repository contract.
type Repository interface {
	Create(ctx context.Context, user *User) error
	Get(ctx context.Context, username string) (*User, error)
}
