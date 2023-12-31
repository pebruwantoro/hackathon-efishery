package task

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/repository"
)

type TaskUsecase interface {
	Create(ctx context.Context, req CreateTaskRequest) error
	GetByID(ctx context.Context, req GetByIdRequest) (GetTaskDetailResponse, error)
	Update(ctx context.Context, req UpdateTaskRequest) error
	GetByObjectiveID(ctx context.Context, req GetByObjectiveIdRequest) ([]GetTaskDetailResponse, error)
	GetBySubtaskID(ctx context.Context, req GetBySubtaskIdRequest) ([]GetTaskDetailResponse, error)
	GetByUserID(ctx context.Context, userID int) (resp []*GetTaskResponse, err error)
}

type usecase struct {
	taskRepository      repository.Tasks
	userLevelRepository repository.UserLevels
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (s *usecase) SetTaskRepository(repo repository.Tasks) *usecase {
	s.taskRepository = repo
	return s
}

func (s *usecase) SetUserLeveRepository(repo repository.UserLevels) *usecase {
	s.userLevelRepository = repo
	return s
}

func (s *usecase) Validate() TaskUsecase {
	if s.taskRepository == nil {
		panic("taskRepository is nil")
	}

	if s.userLevelRepository == nil {
		panic("userLevelRepository is nil")
	}

	return s
}
