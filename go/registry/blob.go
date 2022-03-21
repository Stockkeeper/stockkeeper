package registry

import "github.com/google/uuid"

type Blob struct {
	ID           uuid.UUID
	RepositoryID uuid.UUID
	Location     string
	MediaType    string
	Size         uint
	IsOpen       bool
}

func NewBlob(r *Repository) *Blob {
	return &Blob{
		ID:           uuid.New(),
		RepositoryID: r.ID,
		IsOpen:       true,
	}
}
