package states

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
)

type state struct {
	immutable entities.Immutable
	chain     hash.Hash
	previous  hash.Hash
	height    uint
	trx       []transactions.Transaction
}

func createState(
	immutable entities.Immutable,
	chain hash.Hash,
	previous hash.Hash,
	height uint,
	trx []transactions.Transaction,
) State {
	out := state{
		immutable: immutable,
		chain:     chain,
		previous:  previous,
		height:    height,
		trx:       trx,
	}

	return &out
}

// Hash returns the hash
func (obj *state) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Chain returns the chain
func (obj *state) Chain() hash.Hash {
	return obj.chain
}

// Previous returns the previous hash
func (obj *state) Previous() hash.Hash {
	return obj.previous
}

// Height returns the height
func (obj *state) Height() uint {
	return obj.height
}

// Transactions returns the transactions
func (obj *state) Transactions() []transactions.Transaction {
	return obj.trx
}

// CreatedOn returns the creation time
func (obj *state) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
