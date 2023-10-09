package user

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/repository"
)

type UserUsecase interface {
	Create(ctx context.Context, req CreateUserRequest) error
	Login(ctx context.Context, req LoginRequest) (LoginResponse, error)
	GetUserByUUID(ctx context.Context, req GetUserByUUIDRequest) (GetUserDetailResponse, error)
}

type usecase struct {
	userRepository      repository.Users
	roleRepository      repository.Roles
	userPointRepository repository.UserPoints
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (s *usecase) SetUserRepository(repo repository.Users) *usecase {
	s.userRepository = repo
	return s
}

func (s *usecase) SetRoleRepository(repo repository.Roles) *usecase {
	s.roleRepository = repo
	return s
}

func (s *usecase) SetUserPointRepository(repo repository.UserPoints) *usecase {
	s.userPointRepository = repo
	return s
}

func (s *usecase) Validate() UserUsecase {
	if s.userRepository == nil {
		panic("userRepository is nil")
	}
	if s.roleRepository == nil {
		panic("roleRepository is nil")
	}
	if s.userPointRepository == nil {
		panic("userPointRepository is nil")
	}

	return s
}
