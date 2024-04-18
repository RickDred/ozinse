package repository

import (
	"github.com/RickDred/ozinse/internal/models"
	"github.com/RickDred/ozinse/internal/users"
	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) users.UsersRepositoryInterface {
	return &UsersRepository{db}
}

func (r *UsersRepository) GetUser(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UsersRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UsersRepository) DeleteUser(id uint) error {
	var user models.User
	err := r.db.Delete(&user, id).Error
	return err
}
