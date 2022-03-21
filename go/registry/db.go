package registry

import "github.com/google/uuid"

type DB interface {
	GetRepositoryByName(name string) (*Repository, error)
	GetBlobByID(id uuid.UUID) (*Blob, error)
	InsertBlobAndBlobUploadSession(b *Blob, bus *BlobUploadSession) error
	InsertChunkAndUpdateBlobUploadSession(c *Chunk, bus *BlobUploadSession) error
	DeleteBlobUploadSessionAndUpdateBlob(bus *BlobUploadSession) (*Blob, error)
}
