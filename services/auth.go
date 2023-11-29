package services

import (
	"database/sql"
	"fmt"

	"rest-api-go/models"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Context *gin.Context
	DB      *sql.DB
}

func RegisterUser(c *gin.Context, db *sql.DB) {
	c.Header("Content-Type", "application/json")
	var _ models.User
	fmt.Println(c.Request.Body)
}

// func (c *gin.Context) SignInUser(db *sql.DB) error {
// }

// func DeletePost(c *gin.Context, db *sql.DB) error {
// 	return c.
// }
// func EdiPost(c *gin.Context, db *sql.DB) error {
// 	return c.
// }
