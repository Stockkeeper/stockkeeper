package database

import (
	"github.com/Stockkeeper/stockkeeper/go/registry"
	"github.com/google/uuid"
)

type FakeDatabase struct{}

func (*FakeDatabase) GetRepositoryByName(name string) (*registry.Repository, error) {
	return nil, nil
}
func (*FakeDatabase) GetBlobByID(id uuid.UUID) (*registry.Blob, error) {
	return nil, nil
}
func (*FakeDatabase) InsertBlobAndBlobUploadSession(b *registry.Blob, bus *registry.BlobUploadSession) error {
	return nil
}
func (*FakeDatabase) InsertChunkAndUpdateBlobUploadSession(c *registry.Chunk, bus *registry.BlobUploadSession) error {
	return nil
}
func (*FakeDatabase) DeleteBlobUploadSessionAndUpdateBlob(bus *registry.BlobUploadSession) (*registry.Blob, error) {
	return nil, nil
}
