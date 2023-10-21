package controller

import (
	models "github.com/AlmedinNasufi/go-book-app/Models"
	"github.com/AlmedinNasufi/go-book-app/initializer"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	var books []models.Book
	initializer.DB.Preload("Reviews").Find(&books)
	c.JSON(200, books)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err == nil {
		initializer.DB.Create(&book)
		c.JSON(200, book)
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}
