package wallet

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/wallet/bills"
	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/wallet/statements"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Builder represents a wallet builder
type Builder interface {
	Create() Builder
	WithBills(bills []bills.Bill) Builder
	WithStatement(statement statements.Statement) Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Wallet, error)
}

// Wallet represents the wallet
type Wallet interface {
	entities.Immutable
	Name() string
	Statement() statements.Statement
	HasBills() bool
	Bills() []bills.Bill
	HasDescription() bool
	Description() string
}
