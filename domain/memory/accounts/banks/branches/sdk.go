package branches

import (
	"github.com/xmn-services/rod-network/domain/memory/addresses"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Branch represents a bank branch
type Branch interface {
	entities.Immutable
	TransitNumber() string
	Address() addresses.Address
}
