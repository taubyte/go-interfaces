package seer

import (
	"context"

	"github.com/taubyte/go-interfaces/services"
)

type Service interface {
	services.DBService
	services.GitHubAuth
	Resolver() Resolver
	ListNodes() ([]string, error)
}

type Resolver interface {
	LookupTXT(context.Context, string) ([]string, error)
	LookupCNAME(ctx context.Context, host string) (string, error)
}
