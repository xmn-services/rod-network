package streets

import (
	"github.com/xmn-services/rod-network/domain/memory/addresses/streets"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents a street application
type Application interface {
	Retrieve(hash hash.Hash) (streets.Street, error)
	RetrieveByCity(city hash.Hash) ([]streets.Street, error)
	New(name string, zip string, city hash.Hash) error
	Update(hash hash.Hash, update Update) error
	Merge(hashes []hash.Hash) error
	Delete(hash hash.Hash) error
}

// Update represents an update
type Update interface {
	HasName() bool
	Name() string
	HasCity() bool
	City() hash.Hash
	HasZip() bool
	Zip() string
}
