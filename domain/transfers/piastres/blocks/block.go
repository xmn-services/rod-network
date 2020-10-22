package blocks

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

type block struct {
	immutable  entities.Immutable
	address    hash.Hash
	trx        hashtree.HashTree
	amount     uint
	additional uint
}

func createBlockFromJSON(ins *jsonBlock) (Block, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	address, err := hashAdapter.FromString(ins.Address)
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
		WithAddress(*address).
		WithTransactions(trx).
		WithAmount(ins.Amount).
		WithAdditional(ins.Additional).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createBlock(
	immutable entities.Immutable,
	address hash.Hash,
	trx hashtree.HashTree,
	amount uint,
	additional uint,
) Block {
	out := block{
		immutable:  immutable,
		address:    address,
		trx:        trx,
		amount:     amount,
		additional: additional,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Address returns the address hash
func (obj *block) Address() hash.Hash {
	return obj.address
}

// Transactions returns the transaction hashtree
func (obj *block) Transactions() hashtree.HashTree {
	return obj.trx
}

// Amount returns the amount
func (obj *block) Amount() uint {
	return obj.amount
}

// Additional returns the additional
func (obj *block) Additional() uint {
	return obj.additional
}

// CreatedOn returns the creation time
func (obj *block) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *block) MarshalJSON() ([]byte, error) {
	ins := createJSONBlockFromBlock(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *block) UnmarshalJSON(data []byte) error {
	ins := new(jsonBlock)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createBlockFromJSON(ins)
	if err != nil {
		return err
	}

	insBlock := pr.(*block)
	obj.immutable = insBlock.immutable
	obj.address = insBlock.address
	obj.trx = insBlock.trx
	obj.amount = insBlock.amount
	obj.additional = insBlock.additional
	return nil
}
