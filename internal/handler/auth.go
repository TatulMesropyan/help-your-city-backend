package controllers

import (
	"fmt"
	model "help-your-city-backend/internal/model"
	service "help-your-city-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	service service.UserService
}

type UserHandlers interface {
	RegisterUserHandler(c *gin.Context)
	SignInUserHandler(c *gin.Context)
	ChangePasswordHandler(c *gin.Context)
	ChangePhoneHandler(c *gin.Context)
}

func NewUserHandler(service service.UserService) *User {
	return &User{service: service}
}
func (u *User) RegisterUserHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.service.RegisterUser(user.FirstName, user.LastName, user.Email, user.Birthday, user.Password, user.Phone)
}
func (u *User) SignInUserHandler(c *gin.Context) {
	fmt.Println("Barev")
	email := c.Query("email")
	password := c.Query("password")

	if email == "" || password == "" {
		c.JSON(422, "missing email or password")
		return
	}
	u.service.SignIn(email, password)
}
func (u *User) ChangePasswordHandler(c *gin.Context) {
	newPassword := c.Query("password")
	oldPassword := c.Query("oldPassword")
	id := c.Param("userId")

	if newPassword == "" || oldPassword == "" {
		c.JSON(422, "incorrect email or password")
		return
	}
	if oldPassword == newPassword {
		c.JSON(422, "old password equal new password")
		return
	}
	u.service.ChangePassword(id, newPassword)
}
func (u *User) ChangePhoneHandler(c *gin.Context) {
	id := c.Param("userId")
	phone := c.Query("phone")
	u.service.ChangePhone(id, phone)
}
