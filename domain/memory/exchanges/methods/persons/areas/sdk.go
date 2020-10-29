package areas

import (
	"github.com/xmn-services/rod-network/domain/memory/addresses"
	"github.com/xmn-services/rod-network/domain/memory/addresses/streets"
	"github.com/xmn-services/rod-network/domain/memory/addresses/streets/cities"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Area represents an area
type Area interface {
	entities.Immutable
	IsAddress() bool
	Address() addresses.Address
	IsStreet() bool
	Street() streets.Street
	IsCity() bool
	City() cities.City
}
