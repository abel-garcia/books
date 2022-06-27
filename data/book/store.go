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
	existsError = "Error book exists"
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

func (store *Store) Create(ctx context.Context, book *domain.Book) (string, error) {
	where := map[string]interface{}{"google_id": book.GoogleID, "user_id": book.UserID}
	exists, err := gormBook.Find(ctx, store.db, where)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		appErr := domainErrors.NewAppError(errors.Wrap(err, selectError), domainErrors.RepositoryError)
		return "", appErr
	}

	if exists != nil {
		err := errors.New(existsError)
		appErr := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.ResourceAlreadyExists)
		return "", appErr
	}

	gormBookSchema := &gormBook.Book{
		GoogleID:  book.GoogleID,
		Author:    book.Author,
		Title:     book.Title,
		Publisher: book.Publisher,
		UserID:    uuid.MustParse(book.UserID),
	}

	bookID, err := gormBook.Create(ctx, store.db, gormBookSchema)
	if err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.RepositoryError)
		return "", appErr
	}

	return bookID.String(), nil
}

func (store *Store) GetAll(ctx context.Context, book *domain.Book) ([]domain.Book, error) {

	wishlists, err := gormBook.FindAll(ctx, store.db, &gormBook.Book{
		Author:    book.Author,
		Publisher: book.Publisher,
		GoogleID:  book.GoogleID,
	}, book.WishlistID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		appErr := domainErrors.NewAppError(errors.Wrap(err, selectError), domainErrors.RepositoryError)
		return nil, appErr
	}

	values := []domain.Book{}
	for _, v := range wishlists {
		values = append(values, domain.Book{
			ID:        v.ID.String(),
			GoogleID:  v.GoogleID,
			Author:    v.Author,
			Title:     v.Title,
			Status:    v.Status,
			Publisher: v.Publisher,
			UserID:    v.UserID.String(),
			CreatedAt: v.CreatedAt.Unix(),
			UpdatedAt: v.CreatedAt.Unix(),
		})
	}

	return values, nil
}

func (store *Store) Delete(ctx context.Context, bookID, userID string) error {
	return nil
}
