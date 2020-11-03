package peers

import (
	"github.com/xmn-services/rod-network/domain/memory/peers/peer"
)

type builder struct {
	peers []peer.Peer
}

func createBuilder() Builder {
	out := builder{
		peers: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithPeers add peers to the builder
func (app *builder) WithPeers(peers []peer.Peer) Builder {
	app.peers = peers
	return app
}

// Now builds a new Peers instance
func (app *builder) Now() (Peers, error) {
	if app.peers == nil {
		app.peers = []peer.Peer{}
	}

	mp := map[string]peer.Peer{}
	for _, onePeer := range app.peers {
		mp[onePeer.String()] = onePeer
	}

	lst := []peer.Peer{}
	for _, onePeer := range mp {
		lst = append(lst, onePeer)
	}

	return createPeers(lst), nil
}
