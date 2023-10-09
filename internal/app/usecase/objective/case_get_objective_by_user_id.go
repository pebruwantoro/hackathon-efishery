package objective

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
)

func (s *usecase) GetObjectiveByID(ctx context.Context, req GetObjectiveByUUIDRequest) (result []*ObjectiveResponse, err error) {
	fmt.Println("masuk sini")

	resultObjUser, err := s.objectiveRepository.GetUserObjectiveByUserID(ctx, req.ID)
	if err != nil {
		return
	}

	objID := make([]int, 0)
	for _, objUser := range resultObjUser {
		objID = append(objID, int(objUser.ObjectiveID))
		fmt.Println("objUser.ObjectiveID: ", objUser.ObjectiveID)
	}

	resultTemp, err := s.objectiveRepository.GetObjectiveByID(ctx, objID)

	var tempObjectives []*ObjectiveResponse

	// create a map to quickly look up modules by their ID
	ObjectiveByID := make(map[int]*ObjectiveResponse, len(resultTemp))
	for _, objective := range tempObjectives {
		ObjectiveByID[objective.ID] = objective
	}

	// iterate over the modules again, adding any modules that have a parent ID to their parent's "submenu" slice
	for _, objective := range tempObjectives {
		if objective.SubobjectiveID != 0 {
			parentObjective := ObjectiveByID[objective.SubobjectiveID]
			parentObjective.Subobjective = append(parentObjective.Subobjective, objective)
		}
	}

	// filter out any modules that have a parent ID, leaving only the top-level modules
	topLevelObjectives := make([]*ObjectiveResponse, 0, len(tempObjectives))
	for _, module := range tempObjectives {
		if module.SubobjectiveID == 0 {
			topLevelObjectives = append(topLevelObjectives, module)
		}
	}

	copier.Copy(&result, &topLevelObjectives)

	return
}
