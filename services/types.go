package services

import (
	p2p "github.com/taubyte/go-interfaces/p2p"
)

type Service interface {
	Node() p2p.Node
	Close()
}
