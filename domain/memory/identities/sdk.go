package identities

import (
	"github.com/xmn-services/rod-network/domain/memory/buckets"
	"github.com/xmn-services/rod-network/domain/memory/identities/lists"
	"github.com/xmn-services/rod-network/domain/memory/peers"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills/owners"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Identity represents the identity
type Identity interface {
	entities.Mutable
	Seed() string
	Name() string
	Root() string
	Buckets() buckets.Buckets
	Piastres() owners.Owners
	Lists() lists.Lists
	Peers() peers.Peers
}
