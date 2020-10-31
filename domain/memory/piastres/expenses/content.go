package expenses

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type content struct {
	immutable entities.Immutable
	amount    uint64
	from      []bills.Bill
	lock      locks.Lock
	remaining locks.Lock
}

func createContentFromJSON(ins *JSONContent) (Content, error) {
	billsAdapter := bills.NewAdapter()

	from := []bills.Bill{}
	for _, oneFrom := range ins.From {
		single, err := billsAdapter.ToBill(oneFrom)
		if err != nil {
			return nil, err
		}

		from = append(from, single)
	}

	locksAdapter := locks.NewAdapter()
	lock, err := locksAdapter.ToLock(ins.Lock)
	if err != nil {
		return nil, err
	}

	builder := NewContentBuilder().Create().
		From(from).
		WithLock(lock).
		WithAmount(ins.Amount).
		CreatedOn(ins.CreatedOn)

	if ins.Remaining != nil {
		remaining, err := locksAdapter.ToLock(ins.Remaining)
		if err != nil {
			return nil, err
		}

		builder.WithRemaining(remaining)
	}

	return builder.Now()
}

func createContent(
	immutable entities.Immutable,
	amount uint64,
	from []bills.Bill,
	lock locks.Lock,
) Content {
	return createContentInternally(immutable, amount, from, lock, nil)
}

func createContentWithRemaining(
	immutable entities.Immutable,
	amount uint64,
	from []bills.Bill,
	lock locks.Lock,
	remaining locks.Lock,
) Content {
	return createContentInternally(immutable, amount, from, lock, remaining)
}

func createContentInternally(
	immutable entities.Immutable,
	amount uint64,
	from []bills.Bill,
	lock locks.Lock,
	remaining locks.Lock,
) Content {
	out := content{
		immutable: immutable,
		amount:    amount,
		from:      from,
		lock:      lock,
		remaining: remaining,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Amount returns the amount
func (obj *content) Amount() uint64 {
	return obj.amount
}

// From returns the from bill
func (obj *content) From() []bills.Bill {
	return obj.from
}

// Lock returns the new lock
func (obj *content) Lock() locks.Lock {
	return obj.lock
}

// HasRemaining returns ture if there is a remaining lock, false otherwise
func (obj *content) HasRemaining() bool {
	return obj.remaining != nil
}

// Remaining returns the remaining lock, if any
func (obj *content) Remaining() locks.Lock {
	return obj.remaining
}

// CreatedOn returns the creation time
func (obj *content) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *content) MarshalJSON() ([]byte, error) {
	ins := createJSONContentFromContent(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *content) UnmarshalJSON(data []byte) error {
	ins := new(JSONContent)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createContentFromJSON(ins)
	if err != nil {
		return err
	}

	insExpense := pr.(*content)
	obj.immutable = insExpense.immutable
	obj.amount = insExpense.amount
	obj.from = insExpense.from
	obj.lock = insExpense.lock
	obj.remaining = insExpense.remaining
	return nil
}
