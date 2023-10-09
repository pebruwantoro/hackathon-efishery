package user

import (
	"context"
	"errors"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/helpers"
)

func (u *usecase) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	user := entity.User{}
	if req.Username != "" {
		dataUser, err := u.userRepository.GetByUsername(ctx, req.Username)
		if err != nil {
			return LoginResponse{}, err
		}
		user = dataUser
	} else {
		dataUser, err := u.userRepository.GetByEmail(ctx, req.Email)
		if err != nil {
			return LoginResponse{}, err
		}
		user = dataUser
	}

	// CHECK PASSWORD
	err := helpers.CheckPassword(req.Password, user.Password)
	if err != nil {
		return LoginResponse{}, errors.New("password invalid")
	}

	token, err := helpers.GenerateJWT(ctx, user)
	if err != nil {
		return LoginResponse{}, err
	}

	if token == "" {
		return LoginResponse{}, errors.New("empty token")
	}

	return LoginResponse{
		Username: user.Username,
		Token:    token,
	}, nil
}
