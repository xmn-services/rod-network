package currencies

import (
	"github.com/xmn-services/rod-network/domain/memory/bills/currencies"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents an address application
type Application interface {
	Retrieve(hash hash.Hash) (currencies.Currency, error)
	RetrieveAll() ([]currencies.Currency, error)
	New(symbol string, precision uint, name string, description string) error
	Update(hash hash.Hash, update Update) error
	Merge(hashes []hash.Hash) error
	Delete(hash hash.Hash) error
}

// Update represents an update
type Update interface {
	HasPrecision() bool
	Precision() uint
	HasSymbol() bool
	Symbol() string
	HasName() bool
	Name() string
	HasDescription() bool
	Description() string
}
