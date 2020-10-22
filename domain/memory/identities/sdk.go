package identities

import (
	"github.com/xmn-services/rod-network/domain/memory/buckets"
	"github.com/xmn-services/rod-network/domain/memory/identities/lists"
	"github.com/xmn-services/rod-network/domain/memory/messages"
	"github.com/xmn-services/rod-network/domain/memory/peers"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills/owners"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Identity represents the identity
type Identity interface {
	entities.Mutable
	Name() string
	Root() string
	Buckets() buckets.Buckets
	Piastres() owners.Owners
	Lists() lists.Lists
	Messages() messages.Messages
	Peers() peers.Peers
}
