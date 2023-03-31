package database

type KV interface {
	Get(key string) ([]byte, error)
	Put(key string, v []byte) error
	Delete(key string) error
	List(prefix string) ([]string, error)
	Close()
	UpdateSize(size uint64)
	Size() (uint64, error)
}
