package book

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Create(ctx context.Context, db *gorm.DB, book *Book) (*uuid.UUID, error) {
	if err := db.Create(book).Error; err != nil {
		return nil, err
	}

	return &book.ID, nil
}

func Find(ctx context.Context, db *gorm.DB, where interface{}) (*Book, error) {
	book := &Book{}
	if err := db.Where(where).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func FindByGoogleID(ctx context.Context, db *gorm.DB, googleID string) (*Book, error) {
	book := &Book{}
	if err := db.Where("google_id = ?", googleID).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func FindAll(ctx context.Context, db *gorm.DB, where *Book, wishListID string) ([]Book, error) {
	book := []Book{}

	if wishListID != "" {
		inWithList := db.Select("book_id").Where("wish_list_id = ?", wishListID).Table("wish_list_books")
		if err := db.Where("id IN (?)", inWithList).Where(where).Find(&book).Error; err != nil {
			return nil, err
		}

		return book, nil
	}

	if err := db.Where(where).Find(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}
