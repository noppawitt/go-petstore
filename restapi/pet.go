package restapi

import (
	"github.com/labstack/echo/v4"
	"github.com/noppawitt/go-petstore/spec"
)

func (h *Handler) FindPets(ctx echo.Context, params spec.FindPetsParams) error {
	panic("not implemented")
}

func (h *Handler) AddPet(ctx echo.Context) error {
	panic("not implemented")
}

func (h *Handler) DeletePet(ctx echo.Context, id int64) error {
	panic("not implemented")
}

func (h *Handler) FindPetByID(ctx echo.Context, id int64) error {
	panic("not implemented")
}
