package entries

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type entry struct {
	immutable   entities.Immutable
	name        string
	trx         []hash.Hash
	description string
}

func createEntryFromJSON(ins *jsonEntry) (Entry, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	trx := []hash.Hash{}
	for _, oneTrx := range ins.Transactions {
		hsh, err := hashAdapter.FromString(oneTrx)
		if err != nil {
			return nil, err
		}

		trx = append(trx, *hsh)
	}

	builder := NewBuilder().Create().
		WithHash(*hsh).
		WithName(ins.Name).
		WithTransactions(trx).
		CreatedOn(ins.CreatedOn)

	if ins.Description != "" {
		builder.WithDescription(ins.Description)
	}

	return builder.Now()
}

func createEntry(
	immutable entities.Immutable,
	name string,
	trx []hash.Hash,
) Entry {
	return createEntryInternally(immutable, name, trx, "")
}

func createEntryWithDescription(
	immutable entities.Immutable,
	name string,
	trx []hash.Hash,
	description string,
) Entry {
	return createEntryInternally(immutable, name, trx, description)
}

func createEntryInternally(
	immutable entities.Immutable,
	name string,
	trx []hash.Hash,
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
func (obj *entry) Transactions() []hash.Hash {
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

// MarshalJSON converts the instance to JSON
func (obj *entry) MarshalJSON() ([]byte, error) {
	ins := createJSONEntryFromEntry(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *entry) UnmarshalJSON(data []byte) error {
	ins := new(jsonEntry)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createEntryFromJSON(ins)
	if err != nil {
		return err
	}

	insEntry := pr.(*entry)
	obj.immutable = insEntry.immutable
	obj.name = insEntry.name
	obj.trx = insEntry.trx
	obj.description = insEntry.description
	return nil
}
