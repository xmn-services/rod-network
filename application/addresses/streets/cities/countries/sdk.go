package countries

import (
	"github.com/xmn-services/rod-network/domain/memory/addresses/streets/cities/countries"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents a country application
type Application interface {
	Retrieve(hash hash.Hash) (countries.Country, error)
	RetrieveAll() ([]countries.Country, error)
	New(name string, code string) error
	Update(hash hash.Hash, update Update) error
	Merge(hashes []hash.Hash) error
	Delete(hash hash.Hash) error
}

// Update represents an update application
type Update interface {
	HasName() bool
	Name() string
	HasCode() string
	Code() string
}
