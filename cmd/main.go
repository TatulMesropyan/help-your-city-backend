package main

import (
	"help-your-city-backend/internal/app"
	"help-your-city-backend/internal/config"
	database "help-your-city-backend/internal/database/pg/user"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	router := gin.Default()
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("Error loading config: %v", err)
	}
	db, err := app.Run(&router.RouterGroup, config)
	database.UserRepo(db)
	if err != nil {
		log.Fatal()
	}
	router.Run()
}
