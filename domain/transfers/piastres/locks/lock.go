package locks

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

type lock struct {
	immutable    entities.Immutable
	shareholders hashtree.HashTree
	treeshold    uint
	amount       uint
}

func createLockFromJSON(ins *jsonLock) (Lock, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	ht, err := hashtree.NewAdapter().FromJSON(ins.ShareHolders)
	if err != nil {
		return nil, err
	}

	shareHolders, err := ht.Leaves().HashTree()
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithHash(*hsh).
		WithShareHolders(shareHolders).
		WithTreeshold(ins.Treeshold).
		WithAmount(ins.Amount).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createLock(
	immutable entities.Immutable,
	shareholders hashtree.HashTree,
	treeshold uint,
	amount uint,
) Lock {
	out := lock{
		immutable:    immutable,
		shareholders: shareholders,
		treeshold:    treeshold,
		amount:       amount,
	}

	return &out
}

// Hash returns the hash
func (obj *lock) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// ShareHolders returns the shareholder's hashtree
func (obj *lock) ShareHolders() hashtree.HashTree {
	return obj.shareholders
}

// Treeshold returns the treeshold
func (obj *lock) Treeshold() uint {
	return obj.treeshold
}

// Amount returns the amount of shareholders
func (obj *lock) Amount() uint {
	return obj.amount
}

// CreatedOn returns the creation time
func (obj *lock) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *lock) MarshalJSON() ([]byte, error) {
	ins := createJSONLockFromLock(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *lock) UnmarshalJSON(data []byte) error {
	ins := new(jsonLock)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createLockFromJSON(ins)
	if err != nil {
		return err
	}

	insLock := pr.(*lock)
	obj.immutable = insLock.immutable
	obj.shareholders = insLock.shareholders
	obj.treeshold = insLock.treeshold
	obj.amount = insLock.amount
	return nil
}
