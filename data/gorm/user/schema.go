package user

import (
	"time"

	"github.com/google/uuid"
)

/*
CREATE TABLE IF NOT EXISTS users(
	id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
	username VARCHAR (50) UNIQUE NOT NULL,
	password TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

*/

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
