package mocks

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/ipfs/go-cid"
	mc "github.com/multiformats/go-multicodec"
	mh "github.com/multiformats/go-multihash"
	"github.com/taubyte/go-interfaces/p2p/peer"
)

var cidPrefix = cid.Prefix{
	Version:  1,
	Codec:    uint64(mc.Raw),
	MhType:   mh.SHA2_256,
	MhLength: -1,
}

func (m *mockNode) add(r io.Reader) (_cid cid.Cid, err error) {
	var data []byte
	if data, err = io.ReadAll(r); err != nil {
		return
	}

	if _cid, err = cidPrefix.Sum(data); err != nil {
		return
	}

	m.lock.Lock()
	m.mapDef[_cid.String()] = data
	m.lock.Unlock()

	return
}

func (m *mockNode) get(_cid string) (peer.ReadSeekCloser, error) {
	m.lock.RLock()
	data, exists := m.mapDef[_cid]
	m.lock.RUnlock()
	if !exists {
		return nil, fmt.Errorf("file cid `%s` does not exist", _cid)
	}

	return &mockReadSeekCloser{
		Buffer: bytes.NewBuffer(data),
	}, nil
}

func (m *mockNode) AddFile(r io.Reader) (string, error) {
	_cid, err := m.add(r)
	if err != nil {
		return "", err
	}

	return _cid.String(), nil
}

func (m *mockNode) AddFileForCid(r io.Reader) (cid.Cid, error) {
	return m.add(r)
}

func (m *mockNode) GetFile(_ctx context.Context, id string) (peer.ReadSeekCloser, error) {
	return m.get(id)
}

func (m *mockNode) GetFileFromCid(_ctx context.Context, cid cid.Cid) (peer.ReadSeekCloser, error) {
	return m.get(cid.String())
}

func (m *mockNode) DeleteFile(id string) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, exists := m.mapDef[id]; !exists {
		return fmt.Errorf("file cid `%s` does not exist", id)
	}

	delete(m.mapDef, id)

	return nil
}

func (m *mockNode) Close() {
	for k := range m.mapDef {
		delete(m.mapDef, k)
	}

}

func (m *mockNode) Context() context.Context {
	return m.context
}

func (m *mockReadSeekCloser) Close() error {
	return nil
}

// func (m *mockNode) BootstrapPeers() []peer.AddrInfo {

// }

// func (m *mockNode) DAG() ipfs.Peer {

// }

// func (m *mockNode) Discovery() discovery.Discovery {

// }

// func (m *mockNode) Done() <-chan struct{} {

// }

// func (m *mockNode) ID() peer.ID {

// }
