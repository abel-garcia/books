package wishlist

import "context"

// Repository represent the User's repository contract.
type Repository interface {
	Create(ctx context.Context, wishlist *WishList) (string, error)
	Get(ctx context.Context, wishlist *WishList) ([]WishList, error)
	Delete(ctx context.Context, wishlistID, UserID string) (string, error)
	AddBook(ctx context.Context, wishlistID, bookID, UserID string) error
	DeleteBook(ctx context.Context, wishlistID, bookID, UserID string) error
}
