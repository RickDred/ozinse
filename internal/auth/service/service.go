package service

import (
	"context"

	"github.com/RickDred/ozinse/internal/auth"
	"github.com/RickDred/ozinse/internal/models"
)

type service struct {
	repo auth.Repo
}

func New(repo auth.Repo) auth.Service {
	return &service{repo}
}

func (s *service) Register(ctx context.Context, user *models.User) (string, error) {
	id, err := s.repo.Insert(ctx, user)
	if err != nil {
		return "", err
	}

	token, err := generateJWT(id)
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

	token, err := generateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
