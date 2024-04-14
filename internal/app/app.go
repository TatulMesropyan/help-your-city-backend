package app

import (
	"database/sql"
	api "help-your-city-backend/internal/api/route"
	config "help-your-city-backend/internal/config"
	database "help-your-city-backend/internal/database/pg"
	userStore "help-your-city-backend/internal/database/pg/user"
	handler "help-your-city-backend/internal/handler"
	service "help-your-city-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Run(routerGroup *gin.RouterGroup, config *config.Config) (*sql.DB, error) {
	db, err := database.ConnectDB(config)
	userRepo := userStore.UserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	api.GetAuthRoutes(routerGroup, userHandler)
	if err != nil {
		log.Fatal().Msgf("Error loading config: %v", err)
		return nil, err
	}
	return db, nil
}
