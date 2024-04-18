package transport

import (
	"errors"

	"github.com/RickDred/ozinse/internal/models"
	"github.com/RickDred/ozinse/internal/users"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	service users.UsersServiceInterface
}

func NewHandler(service users.UsersServiceInterface) users.UsersHandlerInterface {
	return &UsersHandler{
		service: service,
	}
}

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
