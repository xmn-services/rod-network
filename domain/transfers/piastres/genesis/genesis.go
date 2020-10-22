package genesis

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type genesis struct {
	immutable               entities.Immutable
	bill                    hash.Hash
	blockDiffBase           uint
	blockDiffIncreasePerTrx float64
	linkDiff                uint
}

func createGenesisFromJSON(ins *jsonGenesis) (Genesis, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	bill, err := hashAdapter.FromString(ins.Bill)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithHash(*hsh).
		WithBill(*bill).
		WithBlockDifficultyBase(ins.BlockDiffBase).
		WithBlockDifficultyIncreasePerTrx(ins.BlockDiffIncreasePerTrx).
		WithLinkDifficulty(ins.LinkDiff).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createGenesis(
	immutable entities.Immutable,
	bill hash.Hash,
	blockDiffBase uint,
	blockDiffIncreasePerTrx float64,
	linkDiff uint,
) Genesis {
	out := genesis{
		immutable:               immutable,
		bill:                    bill,
		blockDiffBase:           blockDiffBase,
		blockDiffIncreasePerTrx: blockDiffIncreasePerTrx,
		linkDiff:                linkDiff,
	}

	return &out
}

// Hash returns the hash
func (obj *genesis) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Bill returns the bill hash
func (obj *genesis) Bill() hash.Hash {
	return obj.bill
}

// BlockDifficultyBase returns the block difficulty base
func (obj *genesis) BlockDifficultyBase() uint {
	return obj.blockDiffBase
}

// BlockDifficultyIncreasePerTrx returns the block difficulty increase per trx
func (obj *genesis) BlockDifficultyIncreasePerTrx() float64 {
	return obj.blockDiffIncreasePerTrx
}

// LinkDifficulty returns the link difficulty
func (obj *genesis) LinkDifficulty() uint {
	return obj.linkDiff
}

// CreatedOn returns the creation time
func (obj *genesis) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *genesis) MarshalJSON() ([]byte, error) {
	ins := createJSONGenesisFromGenesis(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *genesis) UnmarshalJSON(data []byte) error {
	ins := new(jsonGenesis)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createGenesisFromJSON(ins)
	if err != nil {
		return err
	}

	insGenesis := pr.(*genesis)
	obj.immutable = insGenesis.immutable
	obj.bill = insGenesis.bill
	obj.blockDiffBase = insGenesis.blockDiffBase
	obj.blockDiffIncreasePerTrx = insGenesis.blockDiffIncreasePerTrx
	obj.linkDiff = insGenesis.linkDiff
	return nil
}
