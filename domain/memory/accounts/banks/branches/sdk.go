package branches

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/addresses"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Builder represents a branch builder
type Builder interface {
	Create() Builder
	WithTransitNumber(transitNumber string) Builder
	WithAddress(address addresses.Address) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Branch, error)
}

// Branch represents a bank branch
type Branch interface {
	entities.Immutable
	TransitNumber() string
	Address() addresses.Address
}
