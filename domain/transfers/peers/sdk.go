package peers

import (
	"github.com/xmn-services/rod-network/domain/memory/peers/peer"
)

// Adapter represents a peers adapter
type Adapter interface {
	ToPeers(urls []string) (Peers, error)
	ToURLs(peers Peers) []string
}

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

// Repository represents a peers repository
type Repository interface {
	Retrieve() (Peers, error)
}

// Service represents a peer service
type Service interface {
	Save(peers Peers) error
}
