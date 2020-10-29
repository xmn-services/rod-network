package addresses

import (
	"github.com/xmn-services/rod-network/domain/memory/addresses/streets"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Address represents an address
type Address interface {
	entities.Immutable
	Street() streets.Street
	Civic() string
	HasUnit() bool
	Unit() string
}
