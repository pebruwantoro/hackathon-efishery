package user

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/helpers"
)

func (u *usecase) Update(ctx context.Context, req UpdateUserRequest) (err error) {
	user, err := u.userRepository.GetByID(ctx, req.ID)
	if err != nil {
		return
	}
	user.Username = req.Username
	user.Email = req.Email
	user.Username = req.Username
	user.SetUpdated(req.UpdatedBy)

	if req.Password != "" {
		hashedPassword, err := helpers.HashingPassword(req.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}

	err = u.userRepository.Update(ctx, &user)
	if err != nil {
		return
	}

	return nil
}
