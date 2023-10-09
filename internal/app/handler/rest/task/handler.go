package task

import (
	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/task"
)

type TaskHandler interface {
	Create(c echo.Context) (err error)
	GetByID(c echo.Context) (err error)
	GetByObjectiveID(c echo.Context) (err error)
	GetBySubtaskID(c echo.Context) (err error)
	Update(c echo.Context) (err error)
}

type handler struct {
	taskUsecase task.TaskUsecase
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SetTaskUsecase(usecase task.TaskUsecase) *handler {
	h.taskUsecase = usecase
	return h
}

func (h *handler) Validate() TaskHandler {
	if h.taskUsecase == nil {
		panic("taskUsecase is nil")
	}

	return h
}
