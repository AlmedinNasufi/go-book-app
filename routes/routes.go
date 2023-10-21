package routes

import (
	"github.com/AlmedinNasufi/go-book-app/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/authors", controller.GetAllAuthors)
	r.GET("/books", controller.GetAllBooks)
	r.POST("/authors", controller.CreateAuthor)
	r.POST("/books", controller.CreateBook)
	r.POST("/reviews", controller.CreateReview)
}
