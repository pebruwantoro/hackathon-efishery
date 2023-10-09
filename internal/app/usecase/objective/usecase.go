package objective

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/repository"
)

type ObjectiveUsecase interface {
	GetObjectiveByID(ctx context.Context, req GetObjectiveByUUIDRequest) (result []*ObjectiveResponse, err error)
}

type usecase struct {
	objectiveRepository repository.Objectives
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (s *usecase) SetObjectiveRepository(repo repository.Objectives) *usecase {
	s.objectiveRepository = repo
	return s
}
func (s *usecase) Validate() ObjectiveUsecase {
	if s.objectiveRepository == nil {
		panic("objectiveRepository is nil")
	}

	return s
}
