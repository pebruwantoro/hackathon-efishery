package user

import (
	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/user"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/response"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/validator"
)

func (h *handler) Login(c echo.Context) (err error) {
	ctx := c.Request().Context()

	req := user.LoginRequest{}
	if err = validator.Validate(c, &req); err != nil {
		return
	}

	resp, err := h.userUsecase.Login(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
