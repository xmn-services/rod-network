package peer

import "github.com/xmn-services/rod-network/libs/entities"

// Peer represents a peer
type Peer interface {
	entities.Immutable
	Host() string
	Port() uint
}
