package peers

import (
	"github.com/xmn-services/rod-network/domain/memory/peers/peer"
)

// Peers represents peers
type Peers interface {
	All() []peer.Peer
}
