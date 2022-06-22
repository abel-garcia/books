package wishlist

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Create(ctx context.Context, db *gorm.DB, wishlist *WishList) (*uuid.UUID, error) {
	if err := db.Create(wishlist).Error; err != nil {
		return nil, err
	}

	return &wishlist.ID, nil
}

func Find(ctx context.Context, db *gorm.DB, where interface{}) (*WishList, error) {
	wishlist := &WishList{}
	if err := db.Where(where).First(&wishlist).Error; err != nil {
		return nil, err
	}

	return wishlist, nil
}

func FindAll(ctx context.Context, db *gorm.DB, where interface{}) ([]WishList, error) {
	wishlist := []WishList{}
	if err := db.Where(where).Find(&wishlist).Error; err != nil {
		return nil, err
	}

	return wishlist, nil
}
