package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mayhendrap/gin-restful-api/config"
	"net/http"
	"strings"
	"time"
)

func GenerateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Minute * 5).Unix(),
		"subject": email,
	})

	tokenString, err := token.SignedString([]byte(config.Config("SECRET_KEY")))
	if err != nil {
		fmt.Println("error signing", err)
		fmt.Println(config.Config("SECRET_KEY"))
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")

	if bearerToken == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	tokenString := strings.Split(bearerToken, "Bearer ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		return []byte(config.Config("SECRET_KEY")), nil
	})

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
