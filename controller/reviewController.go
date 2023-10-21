package controller

import (
	models "github.com/AlmedinNasufi/go-book-app/Models"
	"github.com/AlmedinNasufi/go-book-app/initializer"
	"github.com/gin-gonic/gin"
)

func CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.BindJSON(&review); err == nil {
		initializer.DB.Create(&review)
		c.JSON(200, review)
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}
