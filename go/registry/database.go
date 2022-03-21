package registry

import "github.com/google/uuid"

type Database interface {
	GetBlobByID(id uuid.UUID) (*Blob, error)
	GetBlobUploadSessionByRepositoryAndID(r *Repository, id string) (*BlobUploadSession, error)
	GetRepositoryByName(name string) (*Repository, error)
	InsertBlobAndBlobUploadSession(b *Blob, bus *BlobUploadSession) error
	InsertChunkAndUpdateBlobUploadSession(c *Chunk, bus *BlobUploadSession) error
	DeleteBlobUploadSessionAndUpdateBlob(bus *BlobUploadSession) (*Blob, error)
}
