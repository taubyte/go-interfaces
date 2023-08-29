package components

import (
	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-interfaces/services/substrate"
	matcherSpec "github.com/taubyte/go-specs/matcher"
)

type ServiceComponent interface {
	substrate.Service
	CheckTns(MatchDefinition) ([]Serviceable, error)
	Cache() Cache
}

/*
GetOptions defines the parameters of serviceables returned by the Cache.Get() method

Validation: if set asset cid, and config commit are validated of the serviceable are validated

Branch: used by the validation method, if not set spec.DefaultBranch is used, currently
this is the only branch handled by production deployed protocols.

MatchIndex: the required Match Index for a serviceable
if not set then matcherSpec.HighMatch is used
*/
type GetOptions struct {
	Validation bool
	Branch     string
	MatchIndex *matcherSpec.Index
}

type Cache interface {
	Add(serviceable Serviceable, branch string) (Serviceable, error)
	Get(MatchDefinition, GetOptions) ([]Serviceable, error)
	Remove(Serviceable)
	Close()
}

type Serviceable interface {
	Match(MatchDefinition) matcherSpec.Index
	Validate(MatchDefinition) error
	Project() (cid.Cid, error)
	Commit() string
	Matcher() MatchDefinition
	Id() string
	Ready() error
	Close()
	Service() ServiceComponent
}

type MatchDefinition interface {
	String() string
	CachePrefix() string
}
