package cities

import (
	"github.com/xmn-services/rod-network/domain/memory/addresses/streets/cities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents a city application
type Application interface {
	Retrieve(hash hash.Hash) (cities.City, error)
	RetrieveByProvince(province hash.Hash) ([]cities.City, error)
	RetrieveByCountry(country hash.Hash) ([]cities.City, error)
	New(name string, province hash.Hash) error
	NewWithCountry(name string, country hash.Hash) error
	Update(hash hash.Hash, update Update) error
	Merge(hashes []hash.Hash) error
	Delete(hash hash.Hash) error
}

// Update represents an update instance
type Update interface {
	HasName() bool
	Name() string
	HasProvince() bool
	Province() hash.Hash
	HasCountry() bool
	Country() hash.Hash
}
