package storage

import "io"

type FakeStorage struct{}

func (*FakeStorage) Get(key string, destination io.Writer) error {
	return nil
}

func (*FakeStorage) Set(key string, source io.Reader) error {
	return nil
}
