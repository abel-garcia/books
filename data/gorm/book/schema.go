package book

import (
	"time"

	"github.com/google/uuid"
	userSchema "github.com/wackGarcia/books/data/gorm/user"
	wishListSchema "github.com/wackGarcia/books/data/gorm/wishlist"
)

/*
	id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    google_id VARCHAR (50) NOT NULL,
	author VARCHAR (50) NOT NULL,
    title VARCHAR (50) NOT NULL,
    status book_status NOT NULL DEFAULT 'active',
    publisher VARCHAR (50) NOT NULL,
	user_id uuid NOT NULL,
    wishlist_id uuid,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(wishlist_id) REFERENCES wishlists(id) ON DELETE CASCADE
*/

type Book struct {
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primary_key"`
	GoogleID   string
	Author     string
	Title      string
	Status     *string
	Publisher  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     uuid.UUID
	WishListID *uuid.UUID
	User       userSchema.User         `gorm:"foreignKey:UserID"`
	WishList   wishListSchema.WishList `gorm:"foreignKey:WishListID"`
}
