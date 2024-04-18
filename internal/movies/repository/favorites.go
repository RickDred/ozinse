package repository

import (
	"context"

	"github.com/RickDred/ozinse/internal/models"
	"github.com/pkg/errors"
)

func (r *MovieRepository) AddToFavorites(ctx context.Context, user *models.User, movieID uint) error {
	if err := r.db.Model(user).Association("Movies").Find(&user.Movies); err != nil {
		return errors.Wrap(err, "failed to preload user's favorite movies")
	}
	for _, favMovie := range user.Movies {
		if favMovie.ID == movieID {
			return errors.New("movie already exists in user's favorites")
		}
	}

	var movie models.Movie
	if err := r.db.First(&movie, movieID).Error; err != nil {
		return errors.Wrap(err, "failed to find movie")
	}

	if err := r.db.Model(user).Association("Movies").Append(&movie); err != nil {
		return errors.Wrap(err, "failed to add movie to user's favorites")
	}

	return nil
}

func (r *MovieRepository) GetFavorites(ctx context.Context, user *models.User) ([]models.Movie, error) {
	var movies []models.Movie

	err := r.db.Preload("Movies").First(user, user.ID).Error
	if err != nil {
		return nil, err
	}

	for _, m := range user.Movies {
		movies = append(movies, *m)
	}

	return movies, nil
}

func (r *MovieRepository) DeleteFavority(ctx context.Context, user *models.User, movieID uint) error {
	if err := r.db.Model(user).Association("Movies").Delete(movieID); err != nil {
		return errors.Wrap(err, "failed to delete movie from user's favorites")
	}
	return nil
}
