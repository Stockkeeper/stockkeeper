package registry

import "io"

type Service struct {
	db      Database
	storage Storage
}

func NewService(db Database, storage Storage) *Service {
	return &Service{db, storage}
}

func (srv *Service) GetRepositoryByName(name string) (*Repository, error) {
	return srv.db.GetRepositoryByName(name)
}

func (srv *Service) OpenBlobUploadSession(r *Repository) (*BlobUploadSession, error) {
	b := NewBlob(r)
	bus := &BlobUploadSession{
		BlobID: b.ID,
	}
	srv.db.InsertBlobAndBlobUploadSession(b, bus)
	return bus, nil
}

func (srv *Service) GetBlobUploadSessionByRepositoryAndID(r *Repository, id string) (*BlobUploadSession, error) {
	return srv.db.GetBlobUploadSessionByRepositoryAndID(r, id)
}

func (srv *Service) CloseBlobUploadSession(bus *BlobUploadSession) (*Blob, error) {
	return srv.db.DeleteBlobUploadSessionAndUpdateBlob(bus)
}

func (srv *Service) AppendChunk(bus *BlobUploadSession, dataReader io.Reader, size uint, rangeStart uint, rangeEnd uint) (*Chunk, error) {
	key := "my-chunk"
	srv.storage.Set(key, dataReader)

	c := &Chunk{
		BlobID:     bus.BlobID,
		Size:       size,
		RangeStart: rangeStart,
		RangeStop:  rangeEnd,
	}
	if err := srv.db.InsertChunkAndUpdateBlobUploadSession(c, bus); err != nil {
		return nil, err
	}
	return c, nil
}
