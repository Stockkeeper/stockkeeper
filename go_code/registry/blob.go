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

type BlobUploadSession struct {
	ID         uuid.UUID
	BlobID     uuid.UUID
	Location   string
	RangeStart uint
}

type Chunk struct {
	ID         uuid.UUID
	BlobID     uuid.UUID
	Location   string
	Size       uint
	RangeStart uint
	RangeStop  uint
}

func AppendChunk(blobUploadSession *BlobUploadSession, location string, size uint, rangeStart uint, rangeStop uint) *Chunk {
	return &Chunk{
		ID:         uuid.New(),
		BlobID:     blobUploadSession.BlobID,
		Location:   location,
		Size:       size,
		RangeStart: rangeStart,
		RangeStop:  rangeStop,
	}
}
