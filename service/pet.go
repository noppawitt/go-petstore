package service

import (
	"context"

	"github.com/noppawitt/go-petstore/domain"
)

type PetService struct {
	store domain.PetStore
}

func NewPetService(store domain.PetStore) *PetService {
	return &PetService{
		store: store,
	}
}

func (s *PetService) ListPets(ctx context.Context, filter *domain.PetFilter) ([]*domain.Pet, error) {
	return s.store.ListPets(ctx, filter)
}

func (s *PetService) FindPet(ctx context.Context, id string) (*domain.Pet, error) {
	return s.store.FindPet(ctx, id)
}

func (s *PetService) CreatePet(ctx context.Context, pet *domain.Pet) error {
	return s.store.CreatePet(ctx, pet)
}

func (s *PetService) DeletePet(ctx context.Context, id string) error {
	return s.store.DeletePet(ctx, id)
}
