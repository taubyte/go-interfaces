package hoarder

import (
	"github.com/ipfs/go-datastore"
	"github.com/taubyte/go-interfaces/services"
)

type Service interface {
	services.Service
	Datastore() datastore.Batching
}