package streets

import (
	"github.com/xmn-services/rod-network/domain/memory/addresses/streets/cities"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Street represents a street
type Street interface {
	entities.Immutable
	City() cities.City
	Name() string
	Zip() string
}
