package wallets

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Adapter represents a wallet adapter
type Adapter interface {
	ToWallet(js []byte) (Wallet, error)
	ToJSON(wallet Wallet) ([]byte, error)
}

// Builder represents a wallet builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithBills(bills []hash.Hash) Builder
	WithStatement(statement hash.Hash) Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Wallet, error)
}

// Wallet represents the wallet
type Wallet interface {
	entities.Immutable
	Name() string
	Statement() hash.Hash
	HasBills() bool
	Bills() []hash.Hash
	HasDescription() bool
	Description() string
}

// Repository represents a wallet repository
type Repository interface {
	Retrieve(hash hash.Hash) (Wallet, error)
}

// Service represents a wallet service
type Service interface {
	Save(wallet Wallet) error
	Delete(wallet Wallet) error
}
