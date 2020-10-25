package bills

import (
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type bill struct {
	immutable entities.Immutable
	bill      hash.Hash
	pks       []signature.PrivateKey
}

func createBill(
	immutable entities.Immutable,
	bll hash.Hash,
	pks []signature.PrivateKey,
) Bill {
	out := bill{
		immutable: immutable,
		bill:      bll,
		pks:       pks,
	}

	return &out
}

// Hash returns the hash
func (obj *bill) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Bill returns the bill
func (obj *bill) Bill() hash.Hash {
	return obj.bill
}

// PrivateKeys returns the privateKeys
func (obj *bill) PrivateKeys() []signature.PrivateKey {
	return obj.pks
}

// CreatedOn returns the creation time
func (obj *bill) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
