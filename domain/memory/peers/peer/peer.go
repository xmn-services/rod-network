package peer

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type peer struct {
	immutable entities.Immutable
	host      string
	port      uint
	isClear   bool
	isOnion   bool
}

func createPeerWithClear(
	immutable entities.Immutable,
	host string,
	port uint,
) Peer {
	return createPeerInternally(immutable, host, port, true, false)
}

func createPeerWithOnion(
	immutable entities.Immutable,
	host string,
	port uint,
) Peer {
	return createPeerInternally(immutable, host, port, false, true)
}

func createPeerInternally(
	immutable entities.Immutable,
	host string,
	port uint,
	isClear bool,
	isOnion bool,
) Peer {
	out := peer{
		immutable: immutable,
		host:      host,
		port:      port,
		isClear:   isClear,
		isOnion:   isOnion,
	}

	return &out
}

// Hash returns the hash
func (obj *peer) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Host returns the host
func (obj *peer) Host() string {
	return obj.host
}

// Port returns the port
func (obj *peer) Port() uint {
	return obj.port
}

// IsClear returns true if the peer is clear, false otherwise
func (obj *peer) IsClear() bool {
	return obj.isClear
}

// IsOnion returns true if the peer is onion, false otherwise
func (obj *peer) IsOnion() bool {
	return obj.isOnion
}

// CreatedOn returns the creation time
func (obj *peer) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
