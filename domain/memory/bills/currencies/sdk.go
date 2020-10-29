package currencies

import "github.com/xmn-services/rod-network/libs/entities"

// Currency represents a currency
type Currency interface {
	entities.Immutable
	Precision() uint
	Symbol() string
	Name() string
	Description() string
}
