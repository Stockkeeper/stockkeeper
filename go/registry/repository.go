package registry

import "github.com/google/uuid"

type Repository struct {
	ID   uuid.UUID
	Name string
}

func ValidateRepositoryName(name string) error {
	return nil
}
