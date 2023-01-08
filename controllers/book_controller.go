package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mayhendrap/gin-restful-api/db"
	"github.com/mayhendrap/gin-restful-api/models"
	bookRequest "github.com/mayhendrap/gin-restful-api/request/book"
	"net/http"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	db.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var request bookRequest.CreateBookRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := models.Book{Title: request.Title, Author: request.Author}
	db.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
	var book models.Book

	if err := db.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	var book models.Book

	if err := db.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	var request bookRequest.UpdateBookRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&book).Updates(&models.Book{ID: book.ID, Title: request.Title, Author: request.Author})
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Book

	if err := db.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	db.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{})
}
