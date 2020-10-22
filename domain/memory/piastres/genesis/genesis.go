package genesis

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
)

type genesis struct {
	immutable  entities.Immutable
	bill       bills.Bill
	difficulty Difficulty
}

func createGenesisFromJSON(ins *JSONGenesis) (Genesis, error) {
	billsAdapter := bills.NewAdapter()
	bill, err := billsAdapter.ToBill(ins.Bill)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithBill(bill).
		WithBlockDifficultyBase(ins.BlockDifficultyBase).
		WithBlockDifficultyIncreasePerTrx(ins.BlockDifficultyIncreasePerTrx).
		WithLinkDifficulty(ins.LinkDifficulty).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createGenesis(
	immutable entities.Immutable,
	bill bills.Bill,
	difficulty Difficulty,
) Genesis {
	out := genesis{
		immutable:  immutable,
		bill:       bill,
		difficulty: difficulty,
	}

	return &out
}

// Hash returns the hash
func (obj *genesis) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Bill returns the bill
func (obj *genesis) Bill() bills.Bill {
	return obj.bill
}

// Difficulty returns the difficulty
func (obj *genesis) Difficulty() Difficulty {
	return obj.difficulty
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
	ins := new(JSONGenesis)
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
	obj.difficulty = insGenesis.difficulty
	return nil
}
