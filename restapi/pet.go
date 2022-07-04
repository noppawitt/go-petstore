package restapi

import (
	"net/http"

	"github.com/noppawitt/go-petstore/spec"
)

// Returns all pets
// (GET /pets)
func (h *Handler) FindPets(w http.ResponseWriter, r *http.Request, params spec.FindPetsParams) {
	panic("not implemented") // TODO: Implement
}

// Creates a new pet
// (POST /pets)
func (h *Handler) AddPet(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}

// Deletes a pet by ID
// (DELETE /pets/{id})
func (h *Handler) DeletePet(w http.ResponseWriter, r *http.Request, id string) {
	panic("not implemented") // TODO: Implement
}

// Returns a pet by ID
// (GET /pets/{id})
func (h *Handler) FindPetByID(w http.ResponseWriter, r *http.Request, id string) {
	panic("not implemented") // TODO: Implement
}
