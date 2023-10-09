package user

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
)

func (u *usecase) GetUserByID(ctx context.Context, req GetUserByIDRequest) (GetUserDetailResponse, error) {
	// GET ENTITY USER
	user, err := u.userRepository.GetByID(ctx, req.Id)
	if err != nil {
		return GetUserDetailResponse{}, err
	}

	point := entity.UserLevel{}
	if user.RoleId != 1 {
		// GET ENTITY USER POINTS
		userPoint, err := u.userLevelRepository.GetByUserID(ctx, req.Id)
		if err != nil {
			return GetUserDetailResponse{}, err
		}
		point = userPoint
	}

	resp := GetUserDetailResponse{
		User: User{
			Id:       user.ID,
			Name:     user.Name,
			UserName: user.Username,
			Email:    user.Email,
		},
		Point: UserPoint{
			Level:           int(point.LevelID),
			HealthPoint:     point.HealthPoint,
			ExperiencePoint: point.ExperiencePoint,
		},
	}

	return resp, nil
}
