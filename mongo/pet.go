package mongo

import (
	"context"

	"github.com/noppawitt/go-petstore/domain"
)

func (s *Store) ListPets(ctx context.Context, filter *domain.PetFilter) ([]*domain.Pet, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Store) FindPet(ctx context.Context, id string) (*domain.Pet, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Store) CreatePet(ctx context.Context, pet *domain.Pet) error {
	panic("not implemented") // TODO: Implement
}

func (s *Store) DeletePet(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}
