package registry

import "github.com/google/uuid"

type BlobUploadSession struct {
	ID         uuid.UUID
	BlobID     uuid.UUID
	Location   string
	RangeStart uint
}
