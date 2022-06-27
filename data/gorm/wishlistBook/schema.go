package wishlistbook

import (
	"time"

	"github.com/google/uuid"
	bookSchema "github.com/wackGarcia/books/data/gorm/book"
	userSchema "github.com/wackGarcia/books/data/gorm/user"
	wishlistSchema "github.com/wackGarcia/books/data/gorm/wishlist"
)

type WishListBook struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primary_key"`
	UserID     uuid.UUID
	WishListID uuid.UUID
	BookID     uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       userSchema.User         `gorm:"foreignKey:UserID"`
	Book       bookSchema.Book         `gorm:"foreignKey:BookID"`
	WishList   wishlistSchema.WishList `gorm:"foreignKey:WishListID"`
}
