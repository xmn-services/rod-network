package entries

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type entry struct {
	immutable   entities.Immutable
	name        string
	trx         []transactions.Transaction
	description string
}

func createEntry(
	immutable entities.Immutable,
	name string,
	trx []transactions.Transaction,
) Entry {
	return createEntryInternally(immutable, name, trx, "")
}

func createEntryWithDescription(
	immutable entities.Immutable,
	name string,
	trx []transactions.Transaction,
	description string,
) Entry {
	return createEntryInternally(immutable, name, trx, description)
}

func createEntryInternally(
	immutable entities.Immutable,
	name string,
	trx []transactions.Transaction,
	description string,
) Entry {
	out := entry{
		immutable:   immutable,
		name:        name,
		trx:         trx,
		description: description,
	}

	return &out
}

// Hash returns the hash
func (obj *entry) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Name returns the name
func (obj *entry) Name() string {
	return obj.name
}

// Transactions returns the transactions
func (obj *entry) Transactions() []transactions.Transaction {
	return obj.trx
}

// CreatedOn returns the creation time
func (obj *entry) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasDescription returns true if there is a description, false otherwise
func (obj *entry) HasDescription() bool {
	return obj.description != ""
}

// Description returns the description, if any
func (obj *entry) Description() string {
	return obj.description
}
