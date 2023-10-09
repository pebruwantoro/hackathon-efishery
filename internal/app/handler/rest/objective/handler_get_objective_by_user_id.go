package objective

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/objective"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/response"
)

func (h *handler) GetSoloObjectiveByUserID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return
	}

	req := objective.GetObjectiveByUUIDRequest{
		ID: id,
	}

	resp, err := h.objectiveUsecase.GetObjectiveByID(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
