package objective

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/repository"
)

type ObjectiveUsecase interface {
	GetObjectiveByID(ctx context.Context, req GetObjectiveByUUIDRequest) (result []*ObjectiveResponse, err error)
	Create(ctx context.Context, req CreateObjectiveRequest) (err error)
	Update(ctx context.Context, req UpdateObjectiveRequest) (err error)
}

type usecase struct {
	objectiveRepository      repository.Objectives
	objectiveUsersRepository repository.ObjectiveUser
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (s *usecase) SetObjectiveRepository(repo repository.Objectives) *usecase {
	s.objectiveRepository = repo
	return s
}

func (s *usecase) SetObjectiveUserRepository(repo repository.ObjectiveUser) *usecase {
	s.objectiveUsersRepository = repo
	return s
}

func (s *usecase) Validate() ObjectiveUsecase {
	if s.objectiveRepository == nil {
		panic("objectiveRepository is nil")
	}
	if s.objectiveUsersRepository == nil {
		panic("objectiveUsersRepository is nil")
	}

	return s
}
