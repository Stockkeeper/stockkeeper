package registry

import "github.com/google/uuid"

type Database interface {
	GetBlobByID(id uuid.UUID) (*Blob, error)
	GetBlobUploadSessionByRepositoryIDAndID(repositoryID, sessionID uuid.UUID) (*BlobUploadSession, error)
	GetImageManifestByRepositoryIDAndRef(repositoryID uuid.UUID, ref string) (*ImageManifest, error)
	GetRepositoryByName(name string) (*Repository, error)
	InsertBlobAndBlobUploadSession(b *Blob, bus *BlobUploadSession) error
	InsertChunkAndUpdateBlobUploadSession(c *Chunk, bus *BlobUploadSession) error
	DeleteBlobUploadSessionAndUpdateBlob(bus *BlobUploadSession) (*Blob, error)
}
