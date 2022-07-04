package mongo

import (
	"context"

	"github.com/noppawitt/go-petstore/domain"
)

type PetStore struct{}

func (s *PetStore) ListPets(ctx context.Context, filter *domain.PetFilter) ([]*domain.Pet, error) {
	panic("not implemented") // TODO: Implement
}

func (s *PetStore) FindPet(ctx context.Context, id string) (*domain.Pet, error) {
	panic("not implemented") // TODO: Implement
}

func (s *PetStore) CreatePet(ctx context.Context, pet *domain.Pet) error {
	panic("not implemented") // TODO: Implement
}

func (s *PetStore) DeletePet(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}
