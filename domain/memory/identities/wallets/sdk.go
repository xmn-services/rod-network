package wallets

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/bills"
	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/statements"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Wallet represents the wallet
type Wallet interface {
	entities.Immutable
	Name() string
	Bills() []bills.Bill
	Balances() Balances
	Statement() statements.Statement
	HasDescription() bool
	Description() string
}

// Balances represents balances
type Balances interface {
	Details() []Balance
	Amount() uint
	BeginsOn() time.Time
	EndsOn() time.Time
}

// Balance represents a balance
type Balance interface {
	Amount() uint
	ValidOn() time.Time
	Expenses() []expenses.Expense
}
