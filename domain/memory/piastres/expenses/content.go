package expenses

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
)

type content struct {
	immutable entities.Immutable
	amount    uint
	from      bills.Bill
	cancel    locks.Lock
	remaining locks.Lock
}

func createContentFromJSON(ins *JSONContent) (Content, error) {
	billsAdapter := bills.NewAdapter()
	from, err := billsAdapter.ToBill(ins.From)
	if err != nil {
		return nil, err
	}

	locksAdapter := locks.NewAdapter()
	cancel, err := locksAdapter.ToLock(ins.Cancel)
	if err != nil {
		return nil, err
	}

	builder := NewContentBuilder().Create().
		From(from).
		WithCancel(cancel).
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
	amount uint,
	from bills.Bill,
	cancel locks.Lock,
) Content {
	return createContentInternally(immutable, amount, from, cancel, nil)
}

func createContentWithRemaining(
	immutable entities.Immutable,
	amount uint,
	from bills.Bill,
	cancel locks.Lock,
	remaining locks.Lock,
) Content {
	return createContentInternally(immutable, amount, from, cancel, remaining)
}

func createContentInternally(
	immutable entities.Immutable,
	amount uint,
	from bills.Bill,
	cancel locks.Lock,
	remaining locks.Lock,
) Content {
	out := content{
		immutable: immutable,
		amount:    amount,
		from:      from,
		cancel:    cancel,
		remaining: remaining,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Amount returns the amount
func (obj *content) Amount() uint {
	return obj.amount
}

// From returns the from bill
func (obj *content) From() bills.Bill {
	return obj.from
}

// Cancel returns the cancel lock
func (obj *content) Cancel() locks.Lock {
	return obj.cancel
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
	obj.cancel = insExpense.cancel
	obj.remaining = insExpense.remaining
	return nil
}
