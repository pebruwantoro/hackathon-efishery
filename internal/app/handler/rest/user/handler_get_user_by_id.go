package user

import (
	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/user"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/response"
)

func (h *handler) GetUserByID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	req := user.GetUserByUUIDRequest{
		Id: c.Param("id"),
	}

	resp, err := h.userUsecase.GetUserByID(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
