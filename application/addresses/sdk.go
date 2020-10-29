package addresses

import (
	"github.com/xmn-services/rod-network/domain/memory/addresses"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents an address application
type Application interface {
	Retrieve(hash hash.Hash) (addresses.Address, error)
	RetrieveByStreet(street hash.Hash) ([]addresses.Address, error)
	New(street hash.Hash, civic string) error
	NewWithUnit(street hash.Hash, civic string, unit string) error
	Update(hash hash.Hash, update Update) error
	Merge(hashes []hash.Hash) error
	Delete(hash hash.Hash) error
}

// Update represents an update
type Update interface {
	HasStreet() bool
	Street() hash.Hash
	HasCivic() bool
	Civic() string
	HasUnit() bool
	Unit() string
}
