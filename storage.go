package storage

import "io"

type Storage interface {
	Get(path string) (io.ReadCloser, error)
	List(path string, page int, pageSize int) []StorageEntity
	Inspect(path string) (StorageEntity, error)
	Create(path string, stream io.ReadCloser) error
	Remove(path string) error
	Exist(path string) (bool, error)
}

type StorageEntity struct {
	Name  string
	Path  string
	Dir   string
	Size  int64
	IsDir bool
	Hash  string
	URL   string
}
