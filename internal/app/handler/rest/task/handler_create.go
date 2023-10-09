package task

import (
	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/task"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/response"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/validator"
)

func (h *handler) Create(c echo.Context) (err error) {
	ctx := c.Request().Context()

	req := task.CreateTaskRequest{
		CreatedBy: c.Request().Header.Get("User-Email"),
	}
	if err = validator.Validate(c, &req); err != nil {
		return
	}

	err = h.taskUsecase.Create(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, nil)
}
