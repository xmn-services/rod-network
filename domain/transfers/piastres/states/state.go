package states

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

type state struct {
	immutable entities.Immutable
	chain     hash.Hash
	prev      hash.Hash
	height    uint
	trx       hashtree.HashTree
	amount    uint
}

func createStateFromJSON(ins *jsonState) (State, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	chain, err := hashAdapter.FromString(ins.Chain)
	if err != nil {
		return nil, err
	}

	prev, err := hashAdapter.FromString(ins.Prev)
	if err != nil {
		return nil, err
	}

	compact, err := hashtree.NewAdapter().FromJSON(ins.Trx)
	if err != nil {
		return nil, err
	}

	trx, err := compact.Leaves().HashTree()
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithHash(*hsh).
		WithChain(*chain).
		WithPrevious(*prev).
		WithHeight(ins.Height).
		WithTransactions(trx).
		WithAmount(ins.Amount).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createState(
	immutable entities.Immutable,
	chain hash.Hash,
	prev hash.Hash,
	height uint,
	trx hashtree.HashTree,
	amount uint,
) State {
	out := state{
		immutable: immutable,
		chain:     chain,
		prev:      prev,
		height:    height,
		trx:       trx,
		amount:    amount,
	}

	return &out
}

// Hash returns the hash
func (obj *state) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Chain returns the chain hash
func (obj *state) Chain() hash.Hash {
	return obj.chain
}

// Previous returns the previous hash
func (obj *state) Previous() hash.Hash {
	return obj.prev
}

// Height returns the height
func (obj *state) Height() uint {
	return obj.height
}

// Transactions returns the transactions
func (obj *state) Transactions() hashtree.HashTree {
	return obj.trx
}

// Amount returns the amount
func (obj *state) Amount() uint {
	return obj.amount
}

// CreatedOn returns the creation time
func (obj *state) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *state) MarshalJSON() ([]byte, error) {
	ins := createJSONStateFromState(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *state) UnmarshalJSON(data []byte) error {
	ins := new(jsonState)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createStateFromJSON(ins)
	if err != nil {
		return err
	}

	insState := pr.(*state)
	obj.immutable = insState.immutable
	obj.chain = insState.chain
	obj.prev = insState.prev
	obj.height = insState.height
	obj.trx = insState.trx
	obj.amount = insState.amount
	return nil
}
