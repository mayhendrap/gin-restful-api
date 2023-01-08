package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mayhendrap/gin-restful-api/controllers"
	"github.com/mayhendrap/gin-restful-api/db"
	"github.com/mayhendrap/gin-restful-api/middlewares"
)

func main() {
	r := gin.Default()

	db.ConnectDatabase()

	bookRoute := r.Group("/", middlewares.RequireAuth())

	bookRoute.GET("/books", controllers.FindBooks)
	bookRoute.POST("/books", controllers.CreateBook)
	bookRoute.GET("/books/:id", controllers.FindBook)
	bookRoute.PATCH("/books/:id", controllers.UpdateBook)
	bookRoute.DELETE("/books/:id", controllers.DeleteBook)

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	err := r.Run()
	if err != nil {
		return
	}
}
