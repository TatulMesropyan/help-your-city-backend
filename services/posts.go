package services

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func  (c *gin.Context) GetAllPosts (db *sql.DB) gin.HandlerFunc {
	return ""
}

func GetSinglePost(c *gin.Context, db *sql.DB) error {
	return c.
}
func DeletePost(c *gin.Context, db *sql.DB) error {
	return c.
}
func EdiPost(c *gin.Context, db *sql.DB) error {
	return c.
}

