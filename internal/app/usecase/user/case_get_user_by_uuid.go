package user

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
)

func (u *usecase) GetUserByUUID(ctx context.Context, req GetUserByUUIDRequest) (GetUserDetailResponse, error) {
	// GET ENTITY USER
	user, err := u.userRepository.GetByUUID(ctx, req.UUID)
	if err != nil {
		return GetUserDetailResponse{}, err
	}

	point := entity.UserPoint{}
	if user.AccessRole != "ADMIN" {
		// GET ENTITY USER POINTS
		userPoint, err := u.userPointRepository.GetByUserUUID(ctx, req.UUID)
		if err != nil {
			return GetUserDetailResponse{}, err
		}
		point = userPoint
	}

	resp := GetUserDetailResponse{
		User: User{
			UUID:     user.UUID,
			Name:     user.Name,
			UserName: user.Username,
			Email:    user.Email,
		},
		Point: UserPoint{
			Level:    point.Level,
			TotalHp:  point.TotalHp,
			TotalExp: point.TotalExp,
		},
	}

	return resp, nil
}
