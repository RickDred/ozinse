package service

import (
	"errors"

	"github.com/RickDred/ozinse/internal/models"
	"github.com/RickDred/ozinse/internal/users"
)

type UsersService struct {
	repo users.UsersRepositoryInterface
}

func NewService(repo users.UsersRepositoryInterface) users.UsersServiceInterface {
	return &UsersService{
		repo: repo,
	}
}

func (s *UsersService) GetProfile(id uint) (*models.User, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}
	user.CleanPassword()
	return user, nil
}

func (s *UsersService) EditProfile(user *models.User, name string, phone int, birthdate string) error {
	reUser, err := s.repo.GetUser(user.ID)
	if err != nil {
		return err
	}
	reUser.Name = name
	reUser.Phone = phone
	reUser.Birthdate = birthdate
	if err := s.repo.UpdateUser(reUser); err != nil {
		return err
	}
	return nil
}

func (s *UsersService) ChangePassword(user *models.User, oldPassword string, newPassword string) error {
	repUser, err := s.repo.GetUser(user.ID)
	if err != nil {
		return err
	}

	user.Password = newPassword
	if err := user.ValidatePassword(); err != nil {
		return err
	}

	if err := repUser.ComparePassword(oldPassword); err != nil {
		return err
	}

	if oldPassword == newPassword {
		return errors.New("new password is the same as old password")
	}
	repUser.Password = newPassword
	repUser.HashPassword()

	if err := s.repo.UpdateUser(repUser); err != nil {
		return err
	}
	return nil
}
