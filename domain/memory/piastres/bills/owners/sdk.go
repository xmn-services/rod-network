package owners

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills/owners/shareholders"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Owner represents a bill owner
type Owner interface {
	entities.Immutable
	ShareHolders() []shareholders.ShareHolder
	Bill() bills.Bill
	PrivateKey() signature.PrivateKey
}
