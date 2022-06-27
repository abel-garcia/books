package wishlistbook

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Create(ctx context.Context, db *gorm.DB, wishlistBook *WishListBook) (*uuid.UUID, error) {
	if err := db.Create(wishlistBook).Error; err != nil {
		return nil, err
	}

	return &wishlistBook.ID, nil
}

func Find(ctx context.Context, db *gorm.DB, where interface{}) (*WishListBook, error) {
	wishlistBook := &WishListBook{}
	if err := db.Where(where).First(&wishlistBook).Error; err != nil {
		return nil, err
	}

	return wishlistBook, nil
}

func Delete(ctx context.Context, db *gorm.DB, where interface{}) (*uuid.UUID, error) {
	wishlistBook := &WishListBook{}
	if err := db.Where(where).Delete(wishlistBook).Error; err != nil {
		return nil, err
	}

	return &wishlistBook.ID, nil
}
