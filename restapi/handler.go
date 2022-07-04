package restapi

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/noppawitt/go-petstore/domain"
	"github.com/noppawitt/go-petstore/spec"
)

type Handler struct {
	PetService domain.PetService
}

func NewHandler() *Handler {
	return &Handler{}
}

var _ spec.ServerInterface = (*Handler)(nil)

// Ping
// (GET /ping)
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func ptrToValue[T any](ptr *T) T {
	var zeroVal T
	if ptr == nil {
		return zeroVal
	}

	return *ptr
}

func renderError(w http.ResponseWriter, r *http.Request, code int, err error) {
	render.Status(r, code)
	render.JSON(w, r, &spec.Error{Error: err.Error()})
}
