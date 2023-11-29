package main

import (
	"log"
	"rest-api-go/config"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register")
		authRouter.POST("/sign-in")
		authRouter.PUT("/change-password")
		authRouter.PUT("/change-email")
		authRouter.PUT("/change-phone")
	}
	postsRouter := router.Group("/posts")
	{
		postsRouter.POST("/all-posts")
		postsRouter.POST("/submit")
		postsRouter.POST("/read")
	}

}
