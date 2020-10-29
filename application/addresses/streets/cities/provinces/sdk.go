package provinces

import (
	"github.com/xmn-services/rod-network/domain/memory/addresses/streets/cities/provinces"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents a province application
type Application interface {
	Retrieve(hash hash.Hash) (provinces.Province, error)
	RetrieveByCountry(country hash.Hash) ([]provinces.Province, error)
	New(country hash.Hash, name string, code string) error
	Update(hash hash.Hash, update Update) error
	Merge(hashes []hash.Hash) error
	Delete(hash hash.Hash) error
}

// Update represents an update application
type Update interface {
	HasCountry() bool
	Country() hash.Hash
	HasName() bool
	Name() string
	HasCode() string
	Code() string
}
