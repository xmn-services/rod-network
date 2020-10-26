package peers

import (
	"github.com/xmn-services/rod-network/domain/memory/peers/peer"
)

// Builder represents peers builder
type Builder interface {
	Create() Builder
	WithPeers(peers []peer.Peer) Builder
	Now() (Peers, error)
}

// Peers represents peers
type Peers interface {
	All() []peer.Peer
}
