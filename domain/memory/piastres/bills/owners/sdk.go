package owners

import (
	"github.com/xmn-services/rod-network/domain/memory/contacts"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Owner represents a bill owner
type Owner interface {
	entities.Immutable
	ShareHolders() []contacts.Contact
	Bill() bills.Bill
	PrivateKey() signature.PrivateKey
}
