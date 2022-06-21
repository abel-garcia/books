package book

import (
	"context"

	"gorm.io/gorm"
)

func Create(ctx context.Context, db *gorm.DB, book *Book) error {
	return db.Create(book).Error
}

func FindByGoogleID(ctx context.Context, db *gorm.DB, googleID string) (*Book, error) {
	book := &Book{}
	if err := db.Where("google_id = ?", googleID).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}
