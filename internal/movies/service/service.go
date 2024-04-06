// internal/movies/service/movie_service.go

package service

import (
	"context"
	"errors"

	"github.com/RickDred/ozinse/internal/models"
	"github.com/RickDred/ozinse/internal/movies"
)

// MovieService implements the MovieService interface.
type MovieService struct {
	movieRepo movies.MovieRepositoryInterface
}

// NewMovieService creates a new instance of MovieServiceImpl.
func NewMovieService(movieRepo movies.MovieRepositoryInterface) movies.MovieServiceInterface {
	return &MovieService{
		movieRepo: movieRepo,
	}
}

// GetMovieByID returns a movie by its ID.
func (s *MovieService) GetMovieByID(ctx context.Context, id string) (*models.Movie, error) {
	// Retrieve the movie from the repository by ID
	movie, err := s.movieRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if movie == nil {
		return nil, errors.New("movie not found")
	}

	return movie, nil
}

// GetMovies returns a list of movies.
func (s *MovieService) GetMovies(ctx context.Context) ([]models.Movie, error) {
	// Retrieve all movies from the repository
	movies, err := s.movieRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

// SearchMovies searches for movies based on a query string.
func (s *MovieService) SearchMovies(ctx context.Context, query string) ([]models.Movie, error) {
	// Search for movies in the repository based on the query string
	// movies, err := s.movieRepo.Search(ctx, query)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

// AddToFavorites adds a movie to a user's favorites.
func (s *MovieService) AddToFavorites(ctx context.Context, userID string, movieID string) error {
	// Add the movie to the user's favorites in the repository
	// err := s.movieRepo.AddToFavorites(ctx, userID, movieID)
	// if err != nil {
	// 	return err
	// }

	return nil
}
