package service

import (
	"context"

	"github.com/RickDred/ozinse/internal/models"
)

func (s *MovieService) AddToFavorites(ctx context.Context, user *models.User, movieID uint) error {
	err := s.repo.AddToFavorites(ctx, user, movieID)
	if err != nil {
		return err
	}

	return nil
}

func (s *MovieService) GetFavorites(ctx context.Context, user *models.User) ([]models.Movie, error) {
	movies, err := s.repo.GetFavorites(ctx, user)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *MovieService) DeleteFavority(ctx context.Context, user *models.User, movieID uint) error {
	err := s.repo.DeleteFavority(ctx, user, movieID)
	if err != nil {
		return err
	}

	return nil
}
