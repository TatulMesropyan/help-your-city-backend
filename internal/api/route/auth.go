package api

import (
	handlers "help-your-city-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func GetAuthRoutes(router *gin.RouterGroup, authHandler handlers.UserHandlers) {
	const (
		AUTH            = "/auth"
		SIGN_IN         = "/sign-in"
		REGISTER        = "/register"
		CHANGE_PHONE    = "/change-phone/:userId"
		CHANGE_PASSWORD = "/change-password/:userId"
	)

	authRouter := router.Group(AUTH)

	authRouter.POST(REGISTER, authHandler.RegisterUserHandler)
	authRouter.POST(SIGN_IN, authHandler.SignInUserHandler)
	authRouter.PUT(CHANGE_PASSWORD, authHandler.ChangePasswordHandler)
	authRouter.PUT(CHANGE_PHONE, authHandler.ChangePhoneHandler)
}
