package addresses

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/addresses/streets"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Builder represents an address builder
type Builder interface {
	Create() Builder
	WithStreet(street streets.Street) Builder
	WithCivic(civic string) Builder
	WithUnit(unit string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Address, error)
}

// Address represents an address
type Address interface {
	entities.Immutable
	Street() streets.Street
	Civic() string
	HasUnit() bool
	Unit() string
}

// Repository represents an address repository
type Repository interface {
	Retrieve(hash hash.Hash) (Address, error)
}
