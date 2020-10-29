package methods

import (
	"github.com/xmn-services/rod-network/domain/memory/accounts"
	"github.com/xmn-services/rod-network/domain/memory/exchanges/methods/persons"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Method represents accepted payment methods
type Method interface {
	entities.Immutable
	Name() string
	HasPersons() bool
	Persons() []persons.Person
	HasDeposits() bool
	Deposits() []accounts.Account
}
