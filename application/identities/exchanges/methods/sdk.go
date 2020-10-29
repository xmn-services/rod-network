package methods

import (
	"github.com/xmn-services/rod-network/domain/memory/exchanges/methods"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents an exchange method application
type Application interface {
	Retrieve(hash hash.Hash) (methods.Method, error)
	RetrieveAll() ([]methods.Method, error)
	New(name string, inPersons []hash.Hash, depositAccounts []hash.Hash) error
	Update(hash hash.Hash, update Update) error
	Delete(hash hash.Hash) error
}

// Update represents an update instance
type Update interface {
	HasName() bool
	Name() string
	HasInPersons() bool
	InPersons() []hash.Hash
	HasDepositAccounts() bool
	DepositAccounts() []hash.Hash
}
