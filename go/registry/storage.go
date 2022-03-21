package registry

import (
	"io"
)

type Storage interface {
	Get(key string, w io.Writer) error
	Set(key string, r io.Reader) error
}
