package container

import (
	"github.com/opentracing/opentracing-go"
	"go.elastic.co/apm/module/apmot/v2"
	"go.elastic.co/apm/v2"

	"github.com/pebruwantoro/hackathon-efishery/config"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/driver"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/repository"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/objective"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/task"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/user"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/logger"
)

type Container struct {
	Config           config.Config
	Tracer           *apm.Tracer
	UserUsecase      user.UserUsecase
	ObjectiveUsecase objective.ObjectiveUsecase
	TaskUsecase      task.TaskUsecase
}

func Setup() *Container {
	// Load Config
	cfg := config.Load()

	logger.NewLogger(logger.Option{IsEnable: cfg.Logger.IsEnable})

	// Setup Tracer
	tracer := driver.NewElasticAPMTracer(cfg.ElasticAPM)

	// Setup opentracing with elastic apm (for helper that need opentracing)
	opentracing.SetGlobalTracer(apmot.New(apmot.WithTracer(tracer)))

	// Setup Driver
	db, _ := driver.NewPostgreSQLDatabase(cfg.DB)

	// Setup Repository
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	userPointRepo := repository.NewUserPointRepository(db)
	objectiveRepo := repository.NewObjectiveRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	// Setup Usecase
	userUsecase := user.NewUsecase().SetUserRepository(userRepo).
		SetRoleRepository(roleRepo).
		SetUserPointRepository(userPointRepo).
		Validate()

	objectiveUsecase := objective.NewUsecase().SetObjectiveRepository(objectiveRepo).
		Validate()

	taskUsecase := task.NewUsecase().
		SetTaskRepository(taskRepo).
		Validate()

	return &Container{
		Config:           cfg,
		Tracer:           tracer,
		UserUsecase:      userUsecase,
		TaskUsecase:      taskUsecase,
		ObjectiveUsecase: objectiveUsecase,
	}
}
