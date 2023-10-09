package task

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
)

func (u *usecase) GetByUserID(ctx context.Context, userID int) (resp []*GetTaskResponse, err error) {
	resultTemp, err := u.taskRepository.GetByUserID(ctx, userID)

	var tempTasks []*GetTaskResponse

	fmt.Println("resultTemp: ", resultTemp)
	copier.Copy(&tempTasks, &resultTemp)

	// create a map to quickly look up modules by their ID
	TaskByID := make(map[int]*GetTaskResponse, len(resultTemp))
	for _, task := range tempTasks {
		TaskByID[int(task.ID)] = task
	}

	fmt.Println("TaskByID: ", TaskByID)
	// iterate over the modules again, adding any modules that have a parent ID to their parent's "submenu" slice
	for _, task := range tempTasks {
		if task.SubtaskID != 0 {
			parentTask := TaskByID[task.SubtaskID]
			parentTask.Subtask = append(parentTask.Subtask, task)
		}
	}

	// filter out any modules that have a parent ID, leaving only the top-level modules
	topLevelTasks := make([]*GetTaskResponse, 0, len(tempTasks))
	for _, module := range tempTasks {
		if module.SubtaskID == 0 {
			topLevelTasks = append(topLevelTasks, module)
		}
	}

	copier.Copy(&resp, &topLevelTasks)
	return
}
