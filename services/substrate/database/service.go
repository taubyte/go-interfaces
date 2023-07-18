package database

import (
	"context"

	logging "github.com/ipfs/go-log/v2"
	"github.com/taubyte/go-interfaces/services"
	smartOps "github.com/taubyte/go-interfaces/services/substrate/smartops"
)

type Service interface {
	services.Service
	Database(context Context) (Database, error)
	Context() context.Context
	Logger() logging.StandardLogger
	Databases() map[string]Database
	Global(projectID string) (Database, error)
	SmartOps() smartOps.Service
}
