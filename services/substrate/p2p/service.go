package p2p

import (
	"context"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/taubyte/go-interfaces/services/substrate/common"
	"github.com/taubyte/go-interfaces/services/substrate/counters"
	smartOps "github.com/taubyte/go-interfaces/services/substrate/smartops"
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/p2p/streams/command"
	"github.com/taubyte/p2p/streams/command/response"
)

type Command interface {
	Send(ctx context.Context, body map[string]interface{}) (response.Response, error)
	SendTo(ctx context.Context, cid cid.Cid, body map[string]interface{}) (response.Response, error)
}

type Stream interface {
	Listen() (protocol string, err error)
	Command(command string) (Command, error)
	Close()
}

type StreamHandler func(cmd command.Command) (resp response.Response, err error)

type CommandService interface {
	Close()
}

type MatchDefinition struct {
	Project     string
	Application string
	Protocol    string
	Command     string
}

func (m *MatchDefinition) String() string {
	return m.Project + m.Application + m.Protocol + m.Command
}

func (m *MatchDefinition) CachePrefix() string {
	return m.Project
}

type Service interface {
	common.Service
	Stream(ctx context.Context, projectID, applicationID, protocol string) (Stream, error)
	StartStream(name, protocol string, handler StreamHandler) (CommandService, error)
	SmartOps() smartOps.Service
	Counter() counters.Service
	LookupService(matcher *MatchDefinition) (config *structureSpec.Service, application string, err error)
	Discover(ctx context.Context, max int, timeout time.Duration) ([]peer.AddrInfo, error)
}

type Serviceable interface {
	common.Serviceable
	Handle(data command.Command) (time.Time, response.Response, error)
	Name() string
	Close()
	Config() *structureSpec.Function
}

type ServiceResource interface {
	Application() string
	Config() *structureSpec.Service
	Context() context.Context
	Project() (cid.Cid, error)
	SmartOps(smartOps []string) (uint32, error)
	Type() uint32
}
