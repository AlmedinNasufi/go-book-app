package main

import (
	"github.com/AlmedinNasufi/go-book-app/initializer"
	"github.com/AlmedinNasufi/go-book-app/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDb()
	initializer.SyncDatabase()
}

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run() // listen
}
