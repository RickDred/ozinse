package transport

import (
	_ "github.com/RickDred/ozinse/docs"
	"github.com/RickDred/ozinse/internal/models"
	"github.com/gin-gonic/gin"
)

type regCredentials struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	RepeatedPassword string `json:"repeated_password"`
}
type tokenResponse struct {
    Token string `json:"token"`
}
type errorResponse struct {
    Error string `json:"error"`
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the provided email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body regCredentials true "Registration credentials"
// @Success 200 {object} tokenResponse "Success"
// @Failure 400 {object} errorResponse "Error"
// @Router /register [post]
func (t *transport) Register(c *gin.Context) {
	var crd regCredentials
	if err := c.ShouldBindJSON(&crd); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Email:    crd.Email,
		Password: crd.Password,
	}

	token, err := t.service.Register(c, user, crd.RepeatedPassword)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func (t *transport) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	token, err := t.service.Login(c, user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func (t *transport) PasswordRecover(c *gin.Context) {
	var credentials struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Email: credentials.Email,
	}
	isok, err := t.service.PasswordRecover(c.Request.Context(), user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"is success": isok})
}
