package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/helpers"
)

func (u *usecase) Create(ctx context.Context, req CreateUserRequest) (err error) {
	user := entity.User{
		UUID:     uuid.NewString(),
		Email:    req.Email,
		Username: req.Username,
		Name:     req.Name,
	}
	user.SetCreated(req.CreatedBy)
	user.SetUpdated(req.CreatedBy)

	// HASHING PASSWORDS
	hashedPassword, err := helpers.HashingPassword(req.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// CHECK ROLE
	role, err := u.roleRepository.GetByID(ctx, req.Role)
	if err != nil {
		return err
	}
	user.AccessRole = role.Name
	fmt.Println("user", user)
	err = u.userRepository.Create(ctx, &user)
	if err != nil {
		return err
	}
	return
}
