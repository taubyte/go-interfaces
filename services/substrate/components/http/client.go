package http

import (
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/taubyte/p2p/streams/client"
)

type Client interface {
	ProxyStreams(host string, path string, method string) (map[peer.ID]*client.Response, map[peer.ID]error, error)
}
