package p2p

import (
	"context"
	"io"

	ds "github.com/ipfs/go-datastore"
	"github.com/taubyte/go-interfaces/p2p/ipfs"

	"github.com/cskr/pubsub"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/core/discovery"
)

type ReadSeekCloser interface {
	io.ReadSeekCloser
	io.WriterTo
}

type Node interface {
	ID() peer.ID
	Peering() PeeringService
	Peer() host.Host
	Messaging() *pubsub.PubSub
	Store() ds.Batching
	DAG() ipfs.Peer
	Discovery() discovery.Discovery
	Context() context.Context
}

type PeeringService interface {
	Start() error
	Stop() error
	AddPeer(peer.AddrInfo)
	RemovePeer(peer.ID)
}
