package wishlist

import (
	"time"

	"github.com/google/uuid"
	userSchema "github.com/wackGarcia/books/data/gorm/user"
)

type WishList struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primary_key"`
	Name      string
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	User      userSchema.User `gorm:"foreignKey:UserID"`
}
