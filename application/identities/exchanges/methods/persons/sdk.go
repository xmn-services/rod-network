package persons

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/exchanges/methods/persons"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents an in-person exchange method application
type Application interface {
	Retrieve(hash hash.Hash) (persons.Person, error)
	RetrieveAll() ([]persons.Person, error)
	NewToAddress(address hash.Hash, time time.Time, maxDelay time.Duration) error
	NewToStreet(street hash.Hash, time time.Time, maxDelay time.Duration) error
	NewToCity(city hash.Hash, time time.Time, maxDelay time.Duration) error
	Delete(hash hash.Hash) error
}
