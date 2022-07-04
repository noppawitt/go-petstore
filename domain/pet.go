package domain

import (
	"context"
	"errors"
)

var ErrPetNotFound = errors.New("pet not found")

type Pet struct {
	ID   string
	Name string
	Tag  *string
}

type PetService interface {
	ListPets(ctx context.Context, filter *PetFilter) ([]*Pet, error)
	FindPet(ctx context.Context, id string) (*Pet, error)
	CreatePet(ctx context.Context, pet *Pet) error
	DeletePet(ctx context.Context, id string) error
}

type PetStore interface {
	ListPets(ctx context.Context, filter *PetFilter) ([]*Pet, error)
	FindPet(ctx context.Context, id string) (*Pet, error)
	CreatePet(ctx context.Context, pet *Pet) error
	DeletePet(ctx context.Context, id string) error
}

type PetFilter struct {
	Tags  []string
	Limit int
}
