package task

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/task"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/response"
)

func (h *handler) GetByObjectiveID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	req := task.GetByObjectiveIdRequest{
		ObjectiveId: id,
	}

	resp, err := h.taskUsecase.GetByObjectiveID(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
