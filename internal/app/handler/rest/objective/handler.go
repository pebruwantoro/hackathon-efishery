package objective

import (
	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/objective"
)

type ObjectiveHandler interface {
	GetSoloObjectiveByUserID(c echo.Context) (err error)
	Create(c echo.Context) (err error)
	Update(c echo.Context) (err error)
}

type handler struct {
	objectiveUsecase objective.ObjectiveUsecase
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SetObjectiveUsecase(usecase objective.ObjectiveUsecase) *handler {
	h.objectiveUsecase = usecase
	return h
}

func (h *handler) Validate() ObjectiveHandler {
	if h.objectiveUsecase == nil {
		panic("objectiveUsecase is nil")
	}

	return h
}
