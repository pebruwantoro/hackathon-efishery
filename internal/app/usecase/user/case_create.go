package user

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/helpers"
)

func (u *usecase) Create(ctx context.Context, req CreateUserRequest) (err error) {
	user := entity.User{
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
	user.RoleId = role.ID

	err = u.userRepository.Create(ctx, &user)
	if err != nil {
		return err
	}

	// GET USER
	dataUser, err := u.userRepository.GetByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	if user.RoleId != 1 {
		point := entity.UserLevel{
			UserID:          dataUser.ID,
			LevelID:         pkg.INITIATE_LEVEL,
			HealthPoint:     pkg.INITIATE_HP,
			ExperiencePoint: pkg.INITIATE_EXP,
		}
		err = u.userLevelRepository.Create(ctx, &point)
		if err != nil {
			return err
		}
	}
	return
}
