package wishlist

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	gormWishlist "github.com/wackGarcia/books/data/gorm/wishlist"
	domainErrors "github.com/wackGarcia/books/domain/errors"
	domain "github.com/wackGarcia/books/domain/wishlist"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Store struct with psql database instances
type Store struct {
	db *gorm.DB
}

const (
	selectError = "Error finding wishlist"
	createError = "Error creating wishlist"
)

// New constructor to store structure
func New(db *sql.DB) *Store {

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	return &Store{
		db: gormDB,
	}
}

func (store *Store) Create(ctx context.Context, wishlist *domain.WishList) (string, error) {
	where := map[string]interface{}{"name": wishlist.Name, "user_id": wishlist.UserID}
	exists, err := gormWishlist.Find(ctx, store.db, where)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		appErr := domainErrors.NewAppError(errors.Wrap(err, selectError), domainErrors.RepositoryError)
		return "", appErr
	}

	if exists != nil {
		err := errors.New("wishlist exists")
		appErr := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.ResourceAlreadyExists)
		return "", appErr
	}

	wishlistSchema := &gormWishlist.WishList{
		Name:   wishlist.Name,
		UserID: uuid.MustParse(wishlist.UserID),
	}
	id, err := gormWishlist.Create(ctx, store.db, wishlistSchema)
	if err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.RepositoryError)
		return "", appErr
	}

	return id.String(), nil
}

func (store *Store) Get(ctx context.Context, wishlist *domain.WishList) ([]domain.WishList, error) {

	where := map[string]interface{}{"id": wishlist.ID, "user_id": wishlist.UserID}
	if wishlist.ID == "" {
		where = map[string]interface{}{"user_id": wishlist.UserID}
	}

	wishlists, err := gormWishlist.FindAll(ctx, store.db, where)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		appErr := domainErrors.NewAppError(errors.Wrap(err, selectError), domainErrors.RepositoryError)
		return nil, appErr
	}

	values := []domain.WishList{}
	for _, v := range wishlists {
		values = append(values, domain.WishList{
			ID:        v.ID.String(),
			Name:      v.Name,
			UserID:    v.UserID.String(),
			CreatedAt: v.CreatedAt.Unix(),
			UpdatedAt: v.CreatedAt.Unix(),
		})
	}

	return values, nil
}
