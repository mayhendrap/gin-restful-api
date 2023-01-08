package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/mayhendrap/gin-restful-api/utils"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.VerifyJWT(c)
	}
}
