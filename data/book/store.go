package book

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	domainErrors "github.com/wackGarcia/books/domain/errors"

	"github.com/pkg/errors"
	gormBook "github.com/wackGarcia/books/data/gorm/book"
	domain "github.com/wackGarcia/books/domain/book"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	selectError = "Error finding book"
	createError = "Error creating book"
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

func (store *Store) Create(ctx context.Context, book *domain.Book) error {
	exists, err := gormBook.FindByGoogleID(ctx, store.db, book.GoogleID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		appErr := domainErrors.NewAppError(errors.Wrap(err, selectError), domainErrors.RepositoryError)
		return appErr
	}

	if exists != nil {
		err := errors.New("book exists")
		appErr := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.ResourceAlreadyExists)
		return appErr
	}

	gormBookSchema := &gormBook.Book{
		GoogleID:   book.GoogleID,
		Author:     book.Author,
		Title:      book.Title,
		Publisher:  book.Publisher,
		UserID:     uuid.MustParse(book.UserID),
		WishListID: nil, //uuid.MustParse(book.WishlistID),
	}
	if err := gormBook.Create(ctx, store.db, gormBookSchema); err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.RepositoryError)
		return appErr
	}

	return nil
}

func (store *Store) Get(ctx context.Context, bookID string) (*domain.Book, error) {
	return nil, nil
}

func (store *Store) GetAll(ctx context.Context, book *domain.Book) ([]domain.Book, error) {
	return nil, nil
}
