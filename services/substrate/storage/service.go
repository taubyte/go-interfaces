package storage

import (
	"context"
	"io"

	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-interfaces/kvdb"
	peer "github.com/taubyte/go-interfaces/p2p/peer"
	"github.com/taubyte/go-interfaces/services"
	smartOps "github.com/taubyte/go-interfaces/services/substrate/smartops"
	structureSpec "github.com/taubyte/go-specs/structure"
)

type Context struct {
	context.Context
	context.CancelFunc
	ProjectId     string
	ApplicationId string
	Matcher       string
	Config        *structureSpec.Storage
}

type Meta interface {
	Get() (io.ReadSeekCloser, error)
	Cid() cid.Cid
	Version() int
}

type Service interface {
	services.Service

	Storages() map[string]Storage
	Get(context Context) (Storage, error)
	Storage(context Context) (Storage, error)
	Add(r io.Reader) (cid.Cid, error)
	GetFile(cid.Cid) (peer.ReadSeekCloser, error)

	SmartOps() smartOps.Service
}

type Storage interface {
	AddFile(r io.ReadSeeker, name string, replace bool) (int, error)
	DeleteFile(name string, version int) error
	Meta(name string, version int) (Meta, error)
	ListVersions(name string) ([]string, error)
	GetLatestVersion(name string) (int, error)
	List(prefix string) ([]string, error)
	Close()
	Used() (int, error)
	Capacity() int
	Id() string
	Kvdb() kvdb.KVDB
	ContextConfig() Context
	UpdateCapacity(size uint64)
}
