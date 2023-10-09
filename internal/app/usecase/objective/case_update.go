package objective

import (
	"context"
	"errors"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
)

func (u *usecase) Update(ctx context.Context, req UpdateObjectiveRequest) error {
	dataObjective, err := u.objectiveRepository.GetObjectiveByID(ctx, []int{req.Id})
	objective := entity.Objective{}
	if len(dataObjective) > 0 {
		objective = dataObjective[0]
	} else {
		return errors.New("objective not found")
	}
	// MAPPING OBJECT
	objective.Name = req.Name
	objective.Description = req.Description
	objective.Weight = req.Weight
	objective.SetUpdated(req.UpdatedBy)

	// UPDATE OBJECTIVE
	err = u.objectiveRepository.Update(ctx, &objective)
	if err != nil {
		return err
	}

	return nil
}
