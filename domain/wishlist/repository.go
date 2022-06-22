package wishlist

import "context"

// Repository represent the User's repository contract.
type Repository interface {
	Create(ctx context.Context, wishlist *WishList) (string, error)
	Get(ctx context.Context, wishlist *WishList) ([]WishList, error)
}
