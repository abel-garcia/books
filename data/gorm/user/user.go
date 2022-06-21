package user

import (
	"context"
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

// New constructor to store structure
func New(db *sql.DB) *Store {
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	return &Store{
		db: gormDB,
	}
}

func Create(ctx context.Context, db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func FindByName(ctx context.Context, db *gorm.DB, username string) (*User, error) {
	user := &User{}
	if err := db.Where("user_name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func Find(ctx context.Context, db *gorm.DB, query interface{}, args ...interface{}) ([]User, error) {
	users := []User{}
	if err := db.Where(query, args...).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
