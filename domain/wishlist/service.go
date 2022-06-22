package wishlist

import "context"

// BookService represent the User's usecases.
type WishListService interface {
	Create(ctx context.Context, wishlist *WishList) (string, error)
	Get(ctx context.Context, wishlistID *WishList) ([]WishList, error)
}

type Service struct {
	repository WishListService
}

func NewWishListService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) Create(ctx context.Context, wishlist *WishList) (string, error) {
	return service.repository.Create(ctx, wishlist)
}

func (service *Service) Get(ctx context.Context, wishlistID *WishList) ([]WishList, error) {
	return service.repository.Get(ctx, wishlistID)
}
