package user

import (
	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/user"
)

type UserHandler interface {
	Create(c echo.Context) (err error)
	Login(c echo.Context) (err error)
	GetUserByID(c echo.Context) (err error)
	Update(c echo.Context) (err error)
}

type handler struct {
	userUsecase user.UserUsecase
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SetUserUsecase(usecase user.UserUsecase) *handler {
	h.userUsecase = usecase
	return h
}

func (h *handler) Validate() UserHandler {
	if h.userUsecase == nil {
		panic("userUsecase is nil")
	}

	return h
}
