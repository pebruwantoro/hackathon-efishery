package rest

import (
	"github.com/labstack/echo/v4"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/container"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/handler/rest/health_check"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/handler/rest/organization"
)

func SetupRouter(server *echo.Echo, container *container.Container) {
	// inject handler with usecase via container
	healthCheckHandler := health_check.NewHandler().Validate()
	organizationHandler := organization.NewHandler().SetOrganizationUsecase(container.OrganizationUsecase).Validate()

	server.GET("/", healthCheckHandler.Check)

	organization := server.Group("/v1/organizations")
	{
		organization.POST("", organizationHandler.Create)
	}
}
