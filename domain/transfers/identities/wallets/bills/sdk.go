package bills

import (
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Adapter represents a bill adapter
type Adapter interface {
	ToBill(js []byte) (Bill, error)
	ToJSON(bill Bill) ([]byte, error)
}

// Builder represents the bill builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithBill(bill hash.Hash) Builder
	WithPrivateKeys(pks []signature.PrivateKey) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Bill, error)
}

// Bill represents a bill in a wallet
type Bill interface {
	entities.Immutable
	Bill() hash.Hash
	PrivateKeys() []signature.PrivateKey
}

// Repository represents a bill repository
type Repository interface {
	Retrieve(hash hash.Hash) (Bill, error)
}

// Service represents a bill service
type Service interface {
	Save(bill Bill) error
	Delete(bill Bill) error
}
