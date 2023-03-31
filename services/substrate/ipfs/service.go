package ipfs

import (
	"context"
	"io"

	"github.com/ipfs/go-cid"
	p2p "github.com/taubyte/go-interfaces/p2p"
)

type Service interface {
	GetFile(ctx context.Context, cid cid.Cid) (p2p.ReadSeekCloser, error)
	AddFile(r io.Reader) (cid.Cid, error)
}
