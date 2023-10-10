package objective

import (
	"context"
	"time"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
)

func (u *usecase) Create(ctx context.Context, req CreateObjectiveRequest) error {
	// MAPPING OBJECT
	var startDate, endDate, dueDate time.Time
	if req.StartDate != "" {
		startTime, _ := time.Parse(time.RFC3339, req.StartDate)
		startDate = startTime
	}
	if req.EndDate != "" {
		endTime, _ := time.Parse(time.RFC3339, req.EndDate)
		endDate = endTime
	}
	if req.DueDate != "" {
		dueTime, _ := time.Parse(time.RFC3339, req.DueDate)
		dueDate = dueTime
	}

	objective := entity.Objective{
		Name:           req.Name,
		Description:    req.Description,
		Weight:         req.Weight,
		SubobjectiveID: req.SubobjectiveID,
		StartDate:      startDate,
		EndDate:        endDate,
		DueDate:        dueDate,
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
