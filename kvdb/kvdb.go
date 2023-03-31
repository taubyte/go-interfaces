package kvdb

import (
	"context"
)

type KVDB interface {
	// Get will retrieve the key indexed data
	Get(ctx context.Context, key string) ([]byte, error)
	// Put will insert the data, indexed by key
	Put(ctx context.Context, key string, v []byte) error
	// Delete deletes the key and index data
	Delete(ctx context.Context, key string) error
	// List lists all keys with the given prefix
	List(ctx context.Context, prefix string) ([]string, error)
	// ListAsync returns a channel to list to listed keys
	ListAsync(ctx context.Context, prefix string) (chan string, error)
	// Batch creates a Batch interface of the current KVDB
	Batch(ctx context.Context) (Batch, error)
	// Sync syncs the KVDB key values
	Sync(ctx context.Context, key string) error
}

type Batch interface {
	Put(key string, value []byte) error
	Delete(key string) error
	Commit() error
}
