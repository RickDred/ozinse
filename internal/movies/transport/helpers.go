package transport

import (
	"errors"

	"github.com/RickDred/ozinse/internal/models"
	"github.com/gin-gonic/gin"
)

func GetUserFromGin(c *gin.Context) (*models.User, error) {
	user, exists := c.Get("user")
	if !exists {
		return nil, errors.New("not Authorized")
	}

	userModel, ok := user.(*models.User)
	if !ok {
		return nil, errors.New("Forbiden")
	}

	return userModel, nil
}
