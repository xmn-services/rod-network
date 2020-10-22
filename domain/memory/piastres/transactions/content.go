package transactions

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/cancels"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
)

type content struct {
	hash              hash.Hash
	triggersOn        time.Time
	executesOnTrigger bool
	fees              expenses.Expense
	expense           expenses.Expense
	cancel            cancels.Cancel
}

func createContentFromJSON(ins *JSONContent) (Content, error) {
	builder := NewContentBuilder().Create().TriggersOn(ins.TriggersOn)
	if ins.ExecutesOnTrigger {
		builder.ExecutesOnTrigger()
	}

	expenseAdapter := expenses.NewAdapter()
	if ins.Fees != nil {
		fees, err := expenseAdapter.ToExpense(ins.Fees)
		if err != nil {
			return nil, err
		}

		builder.WithFees(fees)
	}

	cancelsAdapter := cancels.NewAdapter()
	if ins.Cancel != nil {
		cancel, err := cancelsAdapter.ToCancel(ins.Cancel)
		if err != nil {
			return nil, err
		}

		builder.WithCancel(cancel)
	}

	if ins.Expense != nil {
		expense, err := expenseAdapter.ToExpense(ins.Expense)
		if err != nil {
			return nil, err
		}

		builder.WithExpense(expense)
	}

	return builder.Now()
}

func createContentWithExpense(
	hash hash.Hash,
	triggersOn time.Time,
	executesOnTrigger bool,
	expense expenses.Expense,
) Content {
	return createContentInternally(hash, triggersOn, executesOnTrigger, expense, nil, nil)
}

func createContentWithExpenseAndFees(
	hash hash.Hash,
	triggersOn time.Time,
	executesOnTrigger bool,
	expense expenses.Expense,
	fees expenses.Expense,
) Content {
	return createContentInternally(hash, triggersOn, executesOnTrigger, expense, nil, fees)
}

func createContentWithCancel(
	hash hash.Hash,
	triggersOn time.Time,
	executesOnTrigger bool,
	cancel cancels.Cancel,
) Content {
	return createContentInternally(hash, triggersOn, executesOnTrigger, nil, cancel, nil)
}

func createContentWithCancelAndFees(
	hash hash.Hash,
	triggersOn time.Time,
	executesOnTrigger bool,
	cancel cancels.Cancel,
	fees expenses.Expense,
) Content {
	return createContentInternally(hash, triggersOn, executesOnTrigger, nil, cancel, fees)
}

func createContentInternally(
	hash hash.Hash,
	triggersOn time.Time,
	executesOnTrigger bool,
	expense expenses.Expense,
	cancel cancels.Cancel,
	fees expenses.Expense,
) Content {
	out := content{
		hash:              hash,
		triggersOn:        triggersOn,
		executesOnTrigger: executesOnTrigger,
		expense:           expense,
		cancel:            cancel,
		fees:              fees,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// TriggersOn returns the triggersOn time
func (obj *content) TriggersOn() time.Time {
	return obj.triggersOn
}

// ExecutesOnTrigger returns true if the trx is executed on trigger
func (obj *content) ExecutesOnTrigger() bool {
	return obj.executesOnTrigger
}

// IsExpense returns true if the trx is an expense, false otherwise
func (obj *content) IsExpense() bool {
	return obj.expense != nil
}

// Expense returns the expense, if any
func (obj *content) Expense() expenses.Expense {
	return obj.expense
}

// IsCancel returns true if the trx is a cancel, false otherwise
func (obj *content) IsCancel() bool {
	return obj.cancel != nil
}

// Cancel returns the cancel, if any
func (obj *content) Cancel() cancels.Cancel {
	return obj.cancel
}

// HasFees returns the fees, if any
func (obj *content) HasFees() bool {
	return obj.fees != nil
}

// Fees returns the fees
func (obj *content) Fees() expenses.Expense {
	return obj.fees
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

	insContent := pr.(*content)
	obj.hash = insContent.hash
	obj.triggersOn = insContent.triggersOn
	obj.executesOnTrigger = insContent.executesOnTrigger
	obj.fees = insContent.fees
	obj.expense = insContent.expense
	obj.cancel = insContent.cancel
	return nil
}
