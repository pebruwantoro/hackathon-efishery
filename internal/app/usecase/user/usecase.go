package user

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/repository"
)

type UserUsecase interface {
	Create(ctx context.Context, req CreateUserRequest) error
	Login(ctx context.Context, req LoginRequest) (LoginResponse, error)
}

type usecase struct {
	userRepository repository.Users
	roleRepository repository.Roles
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

func (s *usecase) Validate() UserUsecase {
	if s.userRepository == nil {
		panic("userRepository is nil")
	}
	if s.roleRepository == nil {
		panic("roleRepository is nil")
	}

	return s
}
