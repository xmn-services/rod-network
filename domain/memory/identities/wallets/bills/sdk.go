package bills

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Bill represents a bill in a wallet
type Bill interface {
	entities.Immutable
	Bill() bills.Bill
	PrivateKeys() []signature.PrivateKey
}
