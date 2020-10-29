package cities

import (
	"github.com/xmn-services/rod-network/domain/memory/addresses/streets/cities/countries"
	"github.com/xmn-services/rod-network/domain/memory/addresses/streets/cities/provinces"
	"github.com/xmn-services/rod-network/libs/entities"
)

// City represents a city
type City interface {
	entities.Immutable
	Name() string
	HasProvince() bool
	Province() provinces.Province
	HasCountry() bool
	Country() countries.Country
}
