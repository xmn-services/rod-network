package expenses

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type content struct {
	immutable entities.Immutable
	from      []bills.Bill
	to        bills.Bill
	remaining bills.Bill
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

	to, err := billsAdapter.ToBill(ins.To)
	if err != nil {
		return nil, err
	}

	builder := NewContentBuilder().Create().
		WithAmount(to.Amount()).
		From(from).
		WithLock(to.Lock()).
		CreatedOn(ins.CreatedOn)

	if ins.Remaining != nil {
		remaining, err := billsAdapter.ToBill(ins.Remaining)
		if err != nil {
			return nil, err
		}

		builder.WithRemaining(remaining.Lock())
	}

	return builder.Now()
}

func createContent(
	immutable entities.Immutable,
	from []bills.Bill,
	to bills.Bill,
) Content {
	return createContentInternally(immutable, from, to, nil)
}

func createContentWithRemaining(
	immutable entities.Immutable,
	from []bills.Bill,
	to bills.Bill,
	remaining bills.Bill,
) Content {
	return createContentInternally(immutable, from, to, remaining)
}

func createContentInternally(
	immutable entities.Immutable,
	from []bills.Bill,
	to bills.Bill,
	remaining bills.Bill,
) Content {
	out := content{
		immutable: immutable,
		from:      from,
		to:        to,
		remaining: remaining,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// From returns the from bill
func (obj *content) From() []bills.Bill {
	return obj.from
}

// To returns the new bill
func (obj *content) To() bills.Bill {
	return obj.to
}

// HasRemaining returns ture if there is a remaining lock, false otherwise
func (obj *content) HasRemaining() bool {
	return obj.remaining != nil
}

// Remaining returns the remaining bill, if any
func (obj *content) Remaining() bills.Bill {
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
	obj.from = insExpense.from
	obj.to = insExpense.to
	obj.remaining = insExpense.remaining
	return nil
}
