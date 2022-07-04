package restapi

import (
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
