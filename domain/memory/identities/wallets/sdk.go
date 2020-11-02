package wallets

import (
	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/wallet"
	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/wallet/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
)

// Builder represents a wallets builder
type Builder interface {
	Create() Builder
	WithWallets(wallets []wallet.Wallet) Builder
	Now() (Wallets, error)
}

// Wallets represents wallets
type Wallets interface {
	All() []wallet.Wallet
	Transact(trx transactions.Transaction) error
	Fetch(amount uint64) ([]bills.Bill, error)
}
