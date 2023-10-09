package objective

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
)

func (u *usecase) Update(ctx context.Context, req UpdateObjectiveRequest) error {
	// MAPPING OBJECT
	objective := entity.Objective{
		Name:        req.Name,
		Description: req.Description,
		Weight:      req.Weight,
	}
	objective.SetUpdated(req.UpdatedBy)

	// UPDATE OBJECTIVE
	err := u.objectiveRepository.Update(ctx, &objective)
	if err != nil {
		return err
	}

	return nil
}
