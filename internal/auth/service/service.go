package service

import (
	"context"
	"strings"

	"github.com/RickDred/ozinse/internal/auth"
	"github.com/RickDred/ozinse/internal/models"
)

type service struct {
	repo auth.RepoInterface
}

func New(repo auth.RepoInterface) auth.ServiceInterface {
	return &service{repo}
}

func (s *service) Register(ctx context.Context, user *models.User, repeatedPassword string) (string, error) {
	user.Standardize()

	repeatedPassword = strings.TrimSpace(repeatedPassword)

	if repeatedPassword != user.Password {
		return "", models.ErrWrongPassword
	}

	if err := user.Validate(user.ValidateEmail, user.ValidatePassword); err != nil {
		return "", err
	}

	if err := user.HashPassword(); err != nil {
		return "", err
	}

	if uu, err := s.repo.GetByEmail(ctx, user.Email); uu != nil || err == nil {
		return "", models.ErrEmailExists
	}

	_, err := s.repo.Insert(ctx, user)
	if err != nil {
		return "", err
	}

	user.CleanPassword()

	token, err := generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) Login(ctx context.Context, input *models.User) (string, error) {
	user, err := s.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}

	if err := user.ComparePassword(input.Password); err != nil {
		return "", err
	}

	user.CleanPassword()

	token, err := generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
