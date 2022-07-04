package restapi

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/noppawitt/go-petstore/domain"
	"github.com/noppawitt/go-petstore/spec"
)

// Returns all pets
// (GET /pets)
func (h *Handler) FindPets(w http.ResponseWriter, r *http.Request, params spec.FindPetsParams) {
	ctx := r.Context()

	filter := &domain.PetFilter{
		Tags:  ptrToValue(params.Tags),
		Limit: ptrToValue(params.Limit),
	}

	pets, err := h.PetService.ListPets(ctx, filter)
	if err != nil {
		renderError(w, r, http.StatusInternalServerError, err)
		return
	}

	render.JSON(w, r, pets)
}

// Creates a new pet
// (POST /pets)
func (h *Handler) AddPet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body spec.AddPetJSONRequestBody

	if err := render.DecodeJSON(r.Body, &body); err != nil {
		renderError(w, r, http.StatusInternalServerError, err)
		return
	}

	pet := &domain.Pet{
		Name: body.Name,
		Tag:  body.Tag,
	}

	if err := h.PetService.CreatePet(ctx, pet); err != nil {
		renderError(w, r, http.StatusInternalServerError, err)
		return
	}

	render.JSON(w, r, pet)
}

// Deletes a pet by ID
// (DELETE /pets/{id})
func (h *Handler) DeletePet(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()

	if err := h.PetService.DeletePet(ctx, id); err != nil {
		renderError(w, r, http.StatusInternalServerError, err)
		return
	}

	render.Status(r, http.StatusNoContent)
}

// Returns a pet by ID
// (GET /pets/{id})
func (h *Handler) FindPetByID(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()

	pet, err := h.PetService.FindPet(ctx, id)
	if err != nil {
		switch err {
		case domain.ErrPetNotFound:
			renderError(w, r, http.StatusNotFound, err)
			return
		default:
			renderError(w, r, http.StatusInternalServerError, err)
			return
		}
	}

	render.JSON(w, r, pet)
}
