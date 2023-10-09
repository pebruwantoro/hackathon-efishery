package user

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/repository"
)

type UserUsecase interface {
	Create(ctx context.Context, req CreateUserRequest) error
	Login(ctx context.Context, req LoginRequest) (LoginResponse, error)
	GetUserByID(ctx context.Context, req GetUserByUUIDRequest) (GetUserDetailResponse, error)
	Update(ctx context.Context, req UpdateUserRequest) error
}

type usecase struct {
	userRepository      repository.Users
	roleRepository      repository.Roles
	userLevelRepository repository.UserLevels
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

func (s *usecase) SetUserPointRepository(repo repository.UserLevels) *usecase {
	s.userLevelRepository = repo
	return s
}

func (s *usecase) Validate() UserUsecase {
	if s.userRepository == nil {
		panic("userRepository is nil")
	}
	if s.roleRepository == nil {
		panic("roleRepository is nil")
	}
	if s.userLevelRepository == nil {
		panic("userLevelRepository is nil")
	}

	return s
}
