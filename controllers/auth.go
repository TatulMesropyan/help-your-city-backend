package controllers

import (
	"database/sql"
	"rest-api-go/services"

	"github.com/gin-gonic/gin"
)

//	func AuthController(router *gin.RouterGroup, db *sql.DB) {
//		auth := router.Group("/auth")
//		auth.POST("/register", services.RegisterUser)
//		auth.POST("/sign-in")
//		auth.PUT("/change-password")
//		auth.PUT("/change-email")
//		auth.PUT("/change-phone")
//	}
func (c *gin.Context) AuthController(db *sql.DB) {

}
