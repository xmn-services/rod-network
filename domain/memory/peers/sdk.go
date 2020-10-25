package peers

import (
	"github.com/xmn-services/rod-network/domain/memory/peers/peer"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Peers represents peers
type Peers interface {
	entities.Mutable
	All() []peer.Peer
}
