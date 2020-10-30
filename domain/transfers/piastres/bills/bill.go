package bills

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type bill struct {
	immutable entities.Immutable
	lock      hash.Hash
	amount    uint64
}

func createBillFromJSON(ins *jsonBill) (Bill, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	lock, err := hashAdapter.FromString(ins.Lock)
	if err != nil {
		return nil, err
	}

	return NewBuilder().Create().
		WithHash(*hsh).
		WithLock(*lock).
		WithAmount(ins.Amount).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createBill(
	immutable entities.Immutable,
	lock hash.Hash,
	amount uint64,
) Bill {
	out := bill{
		immutable: immutable,
		lock:      lock,
		amount:    amount,
	}

	return &out
}

// Hash returns the hash
func (obj *bill) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Lock returns the lock hash
func (obj *bill) Lock() hash.Hash {
	return obj.lock
}

// Amount returns the amount
func (obj *bill) Amount() uint64 {
	return obj.amount
}

// CreatedOn returns the creation time
func (obj *bill) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *bill) MarshalJSON() ([]byte, error) {
	ins := createJSONBillFromBill(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *bill) UnmarshalJSON(data []byte) error {
	ins := new(jsonBill)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createBillFromJSON(ins)
	if err != nil {
		return err
	}

	insBill := pr.(*bill)
	obj.immutable = insBill.immutable
	obj.lock = insBill.lock
	obj.amount = insBill.amount
	return nil
}
