package controller

import (
	models "github.com/AlmedinNasufi/go-book-app/Models"
	"github.com/AlmedinNasufi/go-book-app/initializer"
	"github.com/gin-gonic/gin"
)

func GetAllAuthors(c *gin.Context) {
	var authors []models.Author
	initializer.DB.Preload("Books").Preload("Books.Reviews").Find(&authors)
	c.JSON(200, authors)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.BindJSON(&author); err == nil {
		initializer.DB.Create(&author)
		c.JSON(200, author)
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}
