package transport

import (
	"github.com/RickDred/ozinse/internal/models"
	"github.com/gin-gonic/gin"
)

func (t *transport) Register(c *gin.Context) {
	var credentials struct {
		Email            string `json:"email"`
		Password         string `json:"password"`
		RepeatedPassword string `json:"repeated_password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	token, err := t.service.Register(c, user, credentials.RepeatedPassword)
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
