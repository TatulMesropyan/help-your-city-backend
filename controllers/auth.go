package controllers

import (
	"help-your-city-backend/config"
	"help-your-city-backend/services"

	"github.com/gin-gonic/gin"
)

type App struct {
	DB *config.PostgresStore
}

func NewAuthController(app *App) *App {
	return &App{
		DB: app.DB,
	}
}

func (ac *App) RegisterUserHandler(c *gin.Context) {
	services.RegisterUser(ac.DB, c)
}
func (ac *App) SignInUserHandler(c *gin.Context) {
	services.SignInHandler(ac.DB, c)
}
