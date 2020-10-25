package peer

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
)

// Builder represents a peer builder
type Builder interface {
	Create() Builder
	WithHost(host string) Builder
	WithPort(port uint) Builder
	IsClear() Builder
	IsOnion() Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Peer, error)
}

// Peer represents a peer
type Peer interface {
	entities.Immutable
	Host() string
	Port() uint
	IsClear() bool
	IsOnion() bool
}
