// internal/movies/repository/movie_repository.go

package repository

import (
	"context"
	"errors"

	"github.com/RickDred/ozinse/internal/models"
	"github.com/RickDred/ozinse/internal/movies"
	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) movies.MovieRepositoryInterface {
	return &MovieRepository{db}
}

func (r *MovieRepository) GetByID(ctx context.Context, id string) (*models.Movie, error) {
	return nil, errors.New("not implemented")
}

func (r *MovieRepository) GetAll(ctx context.Context) ([]*models.Movie, error) {
	return nil, errors.New("not implemented")
}

// Search searches for movies in the database based on a query string.
func (r *MovieRepository) Search(ctx context.Context, query string) ([]*models.Movie, error) {
	return nil, errors.New("not implemented")
}

// AddToFavorites adds a movie to a user's favorites in the database.
func (r *MovieRepository) AddToFavorites(ctx context.Context, userID string, movieID string) error {
	return errors.New("not implemented")
}
