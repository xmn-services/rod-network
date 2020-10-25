package balance

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Balance represents a balance
type Balance interface {
	entities.Immutable
	Amount() uint64
	ValidOn() time.Time
	Expenses() []expenses.Expense
}
