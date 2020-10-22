package identities

import (
	"github.com/xmn-services/rod-network/domain/memory/buckets"
	"github.com/xmn-services/rod-network/domain/memory/contacts"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills/owners"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Identity represents the identity
type Identity interface {
	entities.Immutable
	Name() string
	Root() string
	HasBuckets() bool
	Buckets() []buckets.Bucket
	HasPiastres() bool
	Piastres() owners.Owner
	HasContacts() bool
	Contacts() []contacts.Contact
}
