package provinces

import (
	"github.com/xmn-services/rod-network/domain/memory/addresses/streets/cities/countries"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Province represents a province
type Province interface {
	entities.Immutable
	Country() countries.Country
	Name() string
	Code() string
}
