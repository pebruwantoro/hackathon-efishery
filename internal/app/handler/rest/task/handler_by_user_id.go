package task

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/response"
)

func (h *handler) GetByUserID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return
	}

	resp, err := h.taskUsecase.GetByUserID(ctx, id)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
