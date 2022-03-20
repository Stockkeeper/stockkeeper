package registry

import "io"

type Service struct {
	db DB
}

func NewService() *Service {
	return &Service{}
}

func (srv *Service) GetRepositoryByName(name string) (*Repository, error) {
	return &Repository{}, nil
}

func (srv *Service) OpenBlobUploadSession(r *Repository) (*BlobUploadSession, error) {
	b := NewBlob(r)
	bus := &BlobUploadSession{
		BlobID: b.ID,
	}
	srv.db.InsertBlobAndBlobUploadSession(b, bus)
	return bus, nil
}

func (srv *Service) CloseBlobUploadSession(bus *BlobUploadSession) (*Blob, error) {
	return srv.db.DeleteBlobUploadSessionAndUpdateBlob(bus)
}

func (srv *Service) AppendChunk(bus *BlobUploadSession, dataReader io.Reader, size uint, rangeStart uint, rangeEnd uint) (*Chunk, error) {
	// TODO: Upload the `data` to object storage.
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
