package repository

import (
	"context"
	"errors"
	"fmt"

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

func (r *MovieRepository) GetByID(ctx context.Context, id uint) (*models.Movie, error) {
	var movie models.Movie
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&movie)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("movie with id %v not found", id)
		}
		return nil, result.Error
	}
	return &movie, nil
}

func (r *MovieRepository) GetAll(ctx context.Context) ([]models.Movie, error) {
	var movies []models.Movie
	result := r.db.WithContext(ctx).Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}
	return movies, nil
}

func (r *MovieRepository) Insert(ctx context.Context, movie *models.Movie) (*models.Movie, error) {
	if err := r.db.WithContext(ctx).Create(movie).Error; err != nil {
		return nil, err
	}
	return movie, nil
}

func (r *MovieRepository) Update(ctx context.Context, id uint, updatedMovie *models.Movie) (*models.Movie, error) {
	if err := r.db.WithContext(ctx).Model(&models.Movie{}).Where("id = ?", id).Updates(updatedMovie).Error; err != nil {
		return nil, err
	}
	return updatedMovie, nil
}

func (r *MovieRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.Movie{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *MovieRepository) GetAllByCategory(ctx context.Context, category string) ([]models.Movie, error) {
	var movies []models.Movie
	if err := r.db.WithContext(ctx).Joins("JOIN movie_categories ON movies.id = movie_categories.movie_id").
		Joins("JOIN categories ON movie_categories.category_id = categories.id").
		Where("categories.name = ?", category).
		Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *MovieRepository) GetMovieSeries(ctx context.Context, movieID uint) ([]models.Video, error) {
	var videos []models.Video
	if err := r.db.WithContext(ctx).Where("movie_id = ?", movieID).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func (r *MovieRepository) UploadVideo(ctx context.Context, video *models.Video) (*models.Video, error) {
	if err := r.db.WithContext(ctx).Create(video).Error; err != nil {
		return nil, err
	}
	return video, nil
}

func (r *MovieRepository) Search(ctx context.Context, filters models.MoviesFilter) ([]models.Movie, error) {
	var movies []models.Movie

	query := r.db.WithContext(ctx)

	if filters.Title != "" {
		query = query.Where("title LIKE ?", "%"+filters.Title+"%")
	}

	if filters.Genre != "" {
		query = query.Joins("JOIN movie_categories ON movies.id = movie_categories.movie_id").
			Joins("JOIN categories ON movie_categories.category_id = categories.id").
			Where("categories.name = ?", filters.Genre)
	}

	if filters.Year != "" {
		query = query.Where("year = ?", filters.Year)
	}

	if filters.Type != "" {
		query = query.Where("type = ?", filters.Type)
	}

	if filters.SortBy != "" {
		switch filters.SortBy {
		case "year":
			if filters.SortDesc {
				query = query.Order("year DESC")
			} else {
				query = query.Order("year ASC")
			}
		case "title":
			if filters.SortDesc {
				query = query.Order("title DESC")
			} else {
				query = query.Order("title ASC")
			}
		}
	}

	if err := query.Find(&movies).Error; err != nil {
		return nil, err
	}

	return movies, nil
}

// wait a minute

func (r *MovieRepository) AddToFavorites(ctx context.Context, user *models.User, movie *models.Movie) error {
	var existingMovie models.Movie
	if err := r.db.Model(user).Where("id = ?", movie.ID).Association("Favorites").Find(&existingMovie); err == nil {
		// Movie already exists in favorites
		return nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// An unexpected error occurred
		return err
	}

	// Add movie to favorites
	if err := r.db.Model(user).Association("Favorites").Append(movie); err != nil {
		return err
	}

	return nil
}
