package user

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/user"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/response"
)

func (h *handler) GetUserByID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	req := user.GetUserByIDRequest{
		Id: id,
	}

	resp, err := h.userUsecase.GetUserByID(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
