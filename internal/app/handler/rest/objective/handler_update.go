package objective

import (
	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/objective"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/response"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/validator"
)

func (h *handler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()

	req := objective.UpdateObjectiveRequest{
		UpdatedBy: c.Request().Header.Get("User-Email"),
	}
	if err = validator.Validate(c, &req); err != nil {
		return
	}

	err = h.objectiveUsecase.Update(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, nil)
}
