package movies

import (
	"context"

	"github.com/RickDred/ozinse/internal/models"
	"github.com/gin-gonic/gin"
)

// MovieHandlerInterface defines the methods implemented by the movie handlers.
type MovieHandlerInterface interface {
	GetMovies(*gin.Context)
	GetMovie(*gin.Context)
	CreateMovie(*gin.Context)
	EditMovie(*gin.Context)
	DeleteMovie(*gin.Context)
	GetMovieSeries(*gin.Context)
	GetMoviesByCategory(*gin.Context)

	UploadVideo(*gin.Context)

	SearchMovies(*gin.Context)

	GetFavorites(*gin.Context)
	AddToFavorites(*gin.Context)
	DeleteFavority(*gin.Context)
}

// MovieServiceInterface defines the methods implemented by the movie service.
type MovieServiceInterface interface {
	GetMovieByID(context.Context, *models.User, uint) (*models.Movie, bool, []models.Movie, error)
	GetMovies(context.Context) ([]models.Movie, error)
	CreateMovie(context.Context, *models.Movie) (*models.Movie, error)
	EditMovie(context.Context, uint, *models.Movie) (*models.Movie, error)
	DeleteMovie(context.Context, uint) error
	GetMovieSeries(context.Context, uint) ([]models.Video, error)
	GetMoviesByCategory(context.Context, string) ([]models.Movie, error)

	UploadVideo(context.Context, *models.Video) (*models.Video, error)

	SearchMovies(ctx context.Context, filters models.MoviesFilter) ([]models.Movie, error)

	AddToFavorites(context.Context, *models.User, uint) error
	GetFavorites(context.Context, *models.User) ([]models.Movie, error)
	DeleteFavority(context.Context, *models.User, uint) error
}

// MovieRepositoryInterface defines the methods implemented by the movie repository.
type MovieRepositoryInterface interface {
	GetByID(context.Context, uint) (*models.Movie, error)
	GetAll(context.Context) ([]models.Movie, error)
	Insert(context.Context, *models.Movie) (*models.Movie, error)
	Update(context.Context, uint, *models.Movie) (*models.Movie, error)
	Delete(context.Context, uint) error
	GetAllByCategory(context.Context, string) ([]models.Movie, error)
	GetMoviesByGenres(context.Context, []models.Genre) ([]models.Movie, error)

	GetMovieSeries(ctx context.Context, movieID uint) ([]models.Video, error)
	UploadVideo(ctx context.Context, video *models.Video) (*models.Video, error)

	Search(ctx context.Context, filters models.MoviesFilter) ([]models.Movie, error)

	AddToFavorites(context.Context, *models.User, uint) error
	GetFavorites(context.Context, *models.User) ([]models.Movie, error)
	DeleteFavority(context.Context, *models.User, uint) error
	IsFavorite(context.Context, *models.User, uint) (bool, error)
}
