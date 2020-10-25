package entries

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Builder represents an entry builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithTransactions(trx []transactions.Transaction) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Entry, error)
}

// Entry represents a statement entry
type Entry interface {
	entities.Immutable
	Name() string
	Transactions() []transactions.Transaction
	HasDescription() bool
	Description() string
}
