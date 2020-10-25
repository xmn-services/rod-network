package peers

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/peers/peer"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Builder represents peers builder
type Builder interface {
	Create() Builder
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
