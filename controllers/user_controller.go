package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mayhendrap/gin-restful-api/db"
	"github.com/mayhendrap/gin-restful-api/models"
	userRequest "github.com/mayhendrap/gin-restful-api/request/user"
	"github.com/mayhendrap/gin-restful-api/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	var request userRequest.RegisterUserRequest
	if c.ShouldBindJSON(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}
	user := models.User{Email: request.Email, Password: string(passwordHash)}
	result := db.DB.Create(&user)

	if result.Error != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to save user"})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func Login(c *gin.Context) {
	var request userRequest.LoginUserRequest
	if c.ShouldBindJSON(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	var user models.User
	db.DB.Find(&user, "email = ?", request.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Credentials invalid"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Credentials invalid"})
		return
	}
	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": token})
}
