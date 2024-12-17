package users

import "gorm.io/gorm"

type Repository struct {
	database *gorm.DB
}

func newRepository(db *gorm.DB) *Repository {
	return &Repository{database: db}
}
