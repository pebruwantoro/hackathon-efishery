package objective

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
)

func (u *usecase) Create(ctx context.Context, req CreateObjectiveRequest) error {
	// MAPPING OBJECT
	objective := entity.Objective{
		Name:        req.Name,
		Description: req.Description,
		Weight:      req.Weight,
	}
	objective.SetCreated(req.CreatedBy)
	objective.SetUpdated(req.CreatedBy)

	// CREATE OBJECTIVE
	err := u.objectiveRepository.Create(ctx, &objective)
	if err != nil {
		return err
	}

	// GET OBJECTIVE
	dataObjective, err := u.objectiveRepository.GetObjectiveByName(ctx, req.Name)
	if err != nil {
		return err
	}

	objectiveUser := entity.ObjectiveUser{
		UserID:      req.UserID,
		ObjectiveID: dataObjective.ID,
	}

	// CREATE OBJECTIVE USERS
	err = u.objectiveUsersRepository.Create(ctx, &objectiveUser)
	if err != nil {
		return err
	}

	return nil
}
