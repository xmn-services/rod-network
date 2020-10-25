package peers

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/peers/peer"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Adapter represents a peers adapter
type Adapter interface {
	ToPeers(js []byte) (Peers, error)
	ToJSON(peers Peers) ([]byte, error)
}

// Builder represents peers builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithPeers(peers []peer.Peer) Builder
	CreatedOn(createdOn time.Time) Builder
	LastUpdatedOn(lastUpdatedOn time.Time) Builder
	Now() (Peers, error)
}

// Peers represents peers
type Peers interface {
	entities.Mutable
	All() []peer.Peer
}

// Repository represents a peers repository
type Repository interface {
	Retrieve() (Peers, error)
}

// Service represents a peer service
type Service interface {
	Save(peers Peers) error
	Update(peers Peers) error
}
