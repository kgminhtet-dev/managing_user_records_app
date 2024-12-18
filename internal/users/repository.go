package users

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetAll(start, end int) (*User, error) {
	return nil, nil
}

func (r *Repository) GetById(id string) (*User, error) {
	return nil, nil
}

func (r *Repository) Update(id string, user *User) (*User, error) {
	return nil, nil
}

func (r *Repository) Delete(id string) (*User, error) {
	return nil, nil
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	return nil, nil
}

func newRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
