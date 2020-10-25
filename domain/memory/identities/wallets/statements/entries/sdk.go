package entries

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Entry represents a statement entry
type Entry interface {
	entities.Immutable
	Name() string
	Transactions() []transactions.Transaction
	HasDescription() bool
	Description() string
}
