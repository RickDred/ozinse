package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *MovieHandler) AddToFavorites(c *gin.Context) {
	user, err := GetUserFromGin(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	movieID, err := strconv.Atoi(c.Param("movieID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.movieService.AddToFavorites(c, user, uint(movieID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie added to favorites successfully"})
}

func (h *MovieHandler) GetFavorites(c *gin.Context) {
	user, err := GetUserFromGin(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	movies, err := h.movieService.GetFavorites(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) DeleteFavority(c *gin.Context) {
	user, err := GetUserFromGin(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	movieID, err := strconv.Atoi(c.Param("movieID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.movieService.DeleteFavority(c, user, uint(movieID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie removed from favorites successfully"})
}
