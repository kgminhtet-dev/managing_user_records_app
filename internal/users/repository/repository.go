package repository

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/data"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Create(user *data.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetAll(start, end int) ([]*data.User, error) {
	var users []*data.User

	err := r.db.Limit(end - start).Offset(start).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) GetById(id string) (*data.User, error) {
	var user data.User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) Update(id string, user *data.User) error {
	return r.db.Model(&data.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&data.User{}).Error
}

func (r *Repository) FindByEmail(email string) (*data.User, error) {
	var user data.User
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
