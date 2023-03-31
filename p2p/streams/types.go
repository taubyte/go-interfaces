package streams

import (
	"io"

	"bitbucket.org/taubyte/p2p/streams"
	"github.com/libp2p/go-libp2p/core/network"
)

type Connection interface {
	io.Closer
	network.ConnSecurity
	network.ConnMultiaddrs
}

type Command interface {
	Connection() (streams.Connection, error)
	Encode(io.Writer) error
}
