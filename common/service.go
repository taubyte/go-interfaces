package common

import (
	"github.com/taubyte/go-interfaces/p2p"
)

type Service interface {
	Node() p2p.Node
	Close()
}
