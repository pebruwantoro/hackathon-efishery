package rest

import (
	"github.com/labstack/echo/v4"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/container"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/handler/rest/health_check"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/handler/rest/objective"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/handler/rest/task"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/handler/rest/user"
)

func SetupRouter(server *echo.Echo, container *container.Container) {
	// inject handler with usecase via container
	healthCheckHandler := health_check.NewHandler().Validate()
	userHandler := user.NewHandler().SetUserUsecase(container.UserUsecase).Validate()
	objectiveHandler := objective.NewHandler().SetObjectiveUsecase(container.ObjectiveUsecase).Validate()
	taskHandler := task.NewHandler().SetTaskUsecase(container.TaskUsecase).Validate()

	server.GET("/", healthCheckHandler.Check)

	admin := server.Group("/v1/users/admin")
	{
		admin.POST("", userHandler.Create)
	}

	users := server.Group("/v1/users")
	{
		users.POST("", userHandler.Create, AuthAdminMiddleware(container))
		users.POST("/login", userHandler.Login)
		users.GET("/:id", userHandler.GetUserByID, AuthMiddleware(container))
		users.PUT("/:id", userHandler.Update, AuthMiddleware(container))
	}

	tasks := server.Group("/v1/tasks")
	{
		tasks.GET("/list/:user_id", taskHandler.GetByUserID)
		tasks.POST("", taskHandler.Create, AuthManagerialMiddleware(container))
		tasks.PUT("/:id", taskHandler.Update, AuthManagerialMiddleware(container))
		tasks.GET("/:id", taskHandler.GetByID, AuthManagerialMiddleware(container))
		tasks.GET("/:id/objectives", taskHandler.GetByObjectiveID, AuthManagerialMiddleware(container))
		tasks.GET("/:id/subtasks", taskHandler.GetBySubtaskID, AuthManagerialMiddleware(container))
	}

	objective := server.Group("/v1/objectives")
	{
		objective.GET("/solo/:user_id", objectiveHandler.GetSoloObjectiveByUserID)
		objective.POST("", objectiveHandler.Create, AuthManagerialMiddleware(container))
		objective.PUT("/:id", objectiveHandler.Update, AuthManagerialMiddleware(container))
	}

}
