package container

import (
	"github.com/opentracing/opentracing-go"
	"go.elastic.co/apm/module/apmot/v2"
	"go.elastic.co/apm/v2"

	"github.com/pebruwantoro/hackathon-efishery/config"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/driver"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/repository"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/organization"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/logger"
)

type Container struct {
	Config              config.Config
	Tracer              *apm.Tracer
	OrganizationUsecase organization.OrganizationUsecase
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
	organizationRepository := repository.NewOrganizationRepository(db)

	// Setup Usecase
	organizationUsecase := organization.NewUsecase().
		SetOrganizationRepository(organizationRepository).
		Validate()

	return &Container{
		Config:              cfg,
		Tracer:              tracer,
		OrganizationUsecase: organizationUsecase,
	}
}
