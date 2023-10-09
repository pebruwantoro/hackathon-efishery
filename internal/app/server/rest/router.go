package rest

import (
	"github.com/labstack/echo/v4"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/container"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/handler/rest/health_check"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/handler/rest/user"
)

func SetupRouter(server *echo.Echo, container *container.Container) {
	// inject handler with usecase via container
	healthCheckHandler := health_check.NewHandler().Validate()
	userHandler := user.NewHandler().SetUserUsecase(container.UserUsecase).Validate()

	server.GET("/", healthCheckHandler.Check)

	users := server.Group("/v1/users")
	{
		users.POST("", userHandler.Create)
		users.POST("/login", userHandler.Login)
	}
}
