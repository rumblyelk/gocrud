package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rumblyelk/gocrud/controllers"
	"github.com/rumblyelk/gocrud/initializers"
	"github.com/rumblyelk/gocrud/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)

	r.DELETE("/posts/:id", controllers.PostsDestroy)

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.LogIn)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
