package peers

import (
	"github.com/xmn-services/rod-network/domain/transfers/peers/peer"
)

type peers struct {
	peers []peer.Peer
}

func createPeers(
	lst []peer.Peer,
) Peers {
	out := peers{
		peers: lst,
	}

	return &out
}

// All returns the peers
func (obj *peers) All() []peer.Peer {
	return obj.peers
}
