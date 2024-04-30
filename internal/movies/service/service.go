package service

import (
	"context"
	"errors"

	"github.com/RickDred/ozinse/internal/models"
	"github.com/RickDred/ozinse/internal/movies"
)

type MovieService struct {
	repo movies.MovieRepositoryInterface
}

func NewMovieService(movieRepo movies.MovieRepositoryInterface) movies.MovieServiceInterface {
	return &MovieService{
		repo: movieRepo,
	}
}

func (s *MovieService) GetMovieByID(ctx context.Context, user *models.User, id uint) (*models.Movie, bool, error) {
	movie, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, false, err
	}
	if movie == nil {
		return nil, false, errors.New("movie not found")
	}

	isFav, err := s.repo.IsFavorite(ctx, user, id)
	if err != nil {
		return nil, false, err
	}

	return movie, isFav, nil
}

func (s *MovieService) CreateMovie(ctx context.Context, movie *models.Movie) (*models.Movie, error) {
	// validate movie

	return s.repo.Insert(ctx, movie)
}

func (s *MovieService) EditMovie(ctx context.Context, id uint, updatedMovie *models.Movie) (*models.Movie, error) {
	// validate movie

	return s.repo.Update(ctx, id, updatedMovie)
}

func (s *MovieService) DeleteMovie(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *MovieService) GetMoviesByCategory(ctx context.Context, categoryID string) ([]models.Movie, error) {
	return s.repo.GetAllByCategory(ctx, categoryID)
}

func (s *MovieService) GetMovies(ctx context.Context) ([]models.Movie, error) {
	movies, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *MovieService) GetMovieSeries(ctx context.Context, movieID uint) ([]models.Video, error) {
	return s.repo.GetMovieSeries(ctx, movieID)
}

func (s *MovieService) UploadVideo(ctx context.Context, video *models.Video) (*models.Video, error) {
	return s.repo.UploadVideo(ctx, video)
}

func (s *MovieService) SearchMovies(ctx context.Context, filters models.MoviesFilter) ([]models.Movie, error) {
	return s.repo.Search(ctx, filters)
}

