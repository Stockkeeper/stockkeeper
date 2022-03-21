package registry

import "github.com/google/uuid"

type Chunk struct {
	ID         uuid.UUID
	BlobID     uuid.UUID
	Location   string
	Size       uint
	RangeStart uint
	RangeStop  uint
}
