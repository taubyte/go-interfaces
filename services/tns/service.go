package tns

import (
	"bitbucket.org/taubyte/go-interfaces/services"
	kv "bitbucket.org/taubyte/kvdb/database"
)

type Service interface {
	services.Service
	KV() *kv.KVDatabase
}
