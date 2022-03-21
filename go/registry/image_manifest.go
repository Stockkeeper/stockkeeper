package registry

import "github.com/google/uuid"

type ImageManifest struct {
	ID     uuid.UUID
	BlobID uuid.UUID
}

func ValidateImageManifestRef(s string) error {
	return nil
}
