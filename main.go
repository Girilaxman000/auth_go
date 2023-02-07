package main

import (
	"github.com/Girilaxman000/auth_go/controllers"
	"github.com/Girilaxman000/auth_go/database"
	"github.com/Girilaxman000/auth_go/initializers"
	"github.com/Girilaxman000/auth_go/middleware"
	"github.com/Girilaxman000/auth_go/migrate"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDatabase()
	migrate.SyncDatabase()
}

func main() {
	router := gin.Default()
	router.POST("/sign_up", controllers.SignUp)
	router.POST("/sign_in", controllers.SignIn)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	router.Run()
}
