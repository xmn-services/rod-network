package peer

import (
	"hash"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
)

// Adapter represents a peer adapter
type Adapter interface {
	ToPeer(js []byte) (Peer, error)
	ToJSON(peer Peer) ([]byte, error)
}

// Builder represents a peer builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
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

// Repository represents a peer repository
type Repository interface {
	Retrieve(hash hash.Hash) (Peer, error)
}

// Service represents a peer service
type Service interface {
	Save(peer Peer) error
	Delete(peer Peer) error
}
