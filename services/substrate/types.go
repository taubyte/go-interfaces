package substrate

import (
	"context"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/taubyte/go-interfaces/services"
	"github.com/taubyte/go-interfaces/services/substrate/counters"
	"github.com/taubyte/go-interfaces/services/substrate/smartops"
	"github.com/taubyte/go-interfaces/services/tns"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/p2p/streams/command/response"
)

type Service interface {
	services.HttpService
	// Vm Returns the  VM service attached to the Substrate
	Vm() vm.Service
	// Tns returns the Tns client attached to the Substrate
	Tns() tns.Client
	// Counter returns the counter service attached to the Substrate s
	Counter() CounterService
	// SmartOps returns the smartops service attached to the Substrate
	SmartOps() SmartOpsService
	Orbitals() []vm.Plugin

	Dev() bool
	Verbose() bool
	Context() context.Context
}

type SmartOpsService interface {
	Service
	Run(caller smartops.EventCaller, smartOpIds []string) (uint32, error)
}

type CounterService interface {
	Service
	Push(...*counters.WrappedMetric)
	Implemented() bool
}

type Client interface {
	Has(host, path, method string, threshold int) (map[peer.ID]response.Response, map[peer.ID]error, error)
	Handle(pid peer.ID) (response.Response, error)
}
