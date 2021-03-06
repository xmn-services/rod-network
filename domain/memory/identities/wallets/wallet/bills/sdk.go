package bills

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Builder represents the bill builder
type Builder interface {
	Create() Builder
	WithBill(bill bills.Bill) Builder
	WithPrivateKey(pk signature.PrivateKey) Builder
	WithRingKeys(ringKeys []signature.PublicKey) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Bill, error)
}

// Bill represents a bill in a wallet
type Bill interface {
	entities.Immutable
	Bill() bills.Bill
	PrivateKey() signature.PrivateKey
	RingKeys() []signature.PublicKey
	Sign() (signature.RingSignature, error)
}
