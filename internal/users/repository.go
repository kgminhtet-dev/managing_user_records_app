package users

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetAll(start, end int) ([]*User, error) {
	var users []*User

	err := r.db.Limit(end - start).Offset(start).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) GetById(id string) (*User, error) {
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) Update(id string, user *User) (*User, error) {
	return nil, nil
}

func (r *Repository) Delete(id string) (*User, error) {
	return nil, nil
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func newRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
