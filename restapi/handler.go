package restapi

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/noppawitt/go-petstore/domain"
	"github.com/noppawitt/go-petstore/spec"
)

type Handler struct {
	petService domain.PetService
}

var _ spec.ServerInterface = (*Handler)(nil)

func ptrToValue[T any](ptr *T) T {
	var zeroVal T
	if ptr != nil {
		return zeroVal
	}

	return *ptr
}

func renderError(w http.ResponseWriter, r *http.Request, code int, err error) {
	render.Status(r, code)
	render.JSON(w, r, &spec.Error{Error: err.Error()})
}
