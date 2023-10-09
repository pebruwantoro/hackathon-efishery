package user

import (
	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/user"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/response"
)

func (h *handler) GetUserByUUID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	req := user.GetUserByUUIDRequest{
		UUID: c.Param("uuid"),
	}

	resp, err := h.userUsecase.GetUserByUUID(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
