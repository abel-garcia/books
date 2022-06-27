package wishlist

import "context"

// BookService represent the User's usecases.
type WishListService interface {
	Create(ctx context.Context, wishlist *WishList) (string, error)
	Get(ctx context.Context, wishlistID *WishList) ([]WishList, error)
	Delete(ctx context.Context, wishlistID, UserID string) (string, error)
	AddBook(ctx context.Context, wishlistID, bookID, UserID string) error
	DeleteBook(ctx context.Context, wishlistID, bookID, UserID string) error
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

func (service *Service) Get(ctx context.Context, wishlist *WishList) ([]WishList, error) {
	return service.repository.Get(ctx, wishlist)
}

func (service *Service) Delete(ctx context.Context, wishlistID, UserID string) (string, error) {
	return service.repository.Delete(ctx, wishlistID, UserID)
}

func (service *Service) AddBook(ctx context.Context, wishlistID, bookID, UserID string) error {
	return service.repository.AddBook(ctx, wishlistID, bookID, UserID)
}

func (service *Service) DeleteBook(ctx context.Context, wishlistID, bookID, UserID string) error {
	return service.repository.DeleteBook(ctx, wishlistID, bookID, UserID)
}
