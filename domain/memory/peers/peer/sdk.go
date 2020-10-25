package peer

import "github.com/xmn-services/rod-network/libs/entities"

// Peer represents a peer
type Peer interface {
	entities.Immutable
	Protocol() Protocol
	Host() string
	Port() uint
}

// Protocol represents a peer protocol
type Protocol interface {
	IsClear() bool
	IsOnion() bool
}
