package main

import (
	"help-your-city-backend/config"
	"help-your-city-backend/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	app := &controllers.App{
		DB: db,
	}

	router := gin.Default()

	authController := controllers.NewAuthController(app)
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", authController.RegisterUserHandler)
		authRouter.POST("/sign-in", authController.SignInUserHandler)
		authRouter.PUT("/change-password/:userId")
		authRouter.PUT("/change-email/:userId")
		authRouter.PUT("/change-phone/:userId")
	}
	postsRouter := router.Group("/posts")
	{
		postsRouter.POST("/all-posts")
		postsRouter.POST("/")
		postsRouter.POST("/read")
	}
	router.Run()
}
