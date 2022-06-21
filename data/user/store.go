package user

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	gormUser "github.com/wackGarcia/books/data/gorm/user"
	domainErrors "github.com/wackGarcia/books/domain/errors"
	domain "github.com/wackGarcia/books/domain/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Store struct with psql database instances
type Store struct {
	db *gorm.DB
}

// New constructor to store structure
func New(db *sql.DB) *Store {
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	return &Store{
		db: gormDB,
	}
}

const (
	selectError = "Error selecting user"
	createError = "Error creating user"
)

func (store *Store) Create(ctx context.Context, user *domain.User) error {
	exists, err := gormUser.FindByName(ctx, store.db, user.UserName)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		appErr := domainErrors.NewAppError(errors.Wrap(err, selectError), domainErrors.RepositoryError)
		return appErr
	}

	if exists != nil {
		err := errors.New("user exists")
		appErr := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.ResourceAlreadyExists)
		return appErr
	}

	if err := gormUser.Create(ctx, store.db, &gormUser.User{
		UserName: user.UserName,
		Password: user.Password,
	}); err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.RepositoryError)
		return appErr
	}

	return nil
}

func (store *Store) Get(ctx context.Context, username string) (*domain.User, error) {
	user, err := gormUser.FindByName(ctx, store.db, username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		appErr := domainErrors.NewAppErrorWithType(domainErrors.NotAuthorized)
		return nil, appErr
	}

	if err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, selectError), domainErrors.RepositoryError)
		return nil, appErr
	}

	return &domain.User{
		ID:        user.ID.String(),
		UserName:  user.UserName,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.CreatedAt.Unix(),
	}, nil
}
