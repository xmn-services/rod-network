package transactions

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/cancels"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/libs/hash"
)

type content struct {
	hash       hash.Hash
	triggersOn time.Time
	element    Element
	fees       []expenses.Expense
}

func createContentFromJSON(js *JSONContent) (Content, error) {
	hashAdapter := hash.NewAdapter()
	elementBuilder := NewElementBuilder().Create()
	if js.Bucket != "" {
		bucket, err := hashAdapter.FromString(js.Bucket)
		if err != nil {
			return nil, err
		}

		elementBuilder.WithBucket(*bucket)
	}

	if js.Cancel != nil {
		cancel, err := cancels.NewAdapter().ToCancel(js.Cancel)
		if err != nil {
			return nil, err
		}

		elementBuilder.WithCancel(cancel)
	}

	builder := NewContentBuilder().Create().TriggersOn(js.TriggersOn)
	if len(js.Fees) > 0 {
		expenseAdapter := expenses.NewAdapter()
		fees := []expenses.Expense{}
		for _, oneFee := range js.Fees {
			fee, err := expenseAdapter.ToExpense(oneFee)
			if err != nil {
				return nil, err
			}

			fees = append(fees, fee)
		}

		builder.WithFees(fees)
	}

	if (js.Bucket != "") || (js.Cancel != nil) {
		element, err := elementBuilder.Now()
		if err != nil {
			return nil, err
		}

		builder.WithElement(element)
	}

	return builder.Now()
}

func createContentWithElement(
	hash hash.Hash,
	triggersOn time.Time,
	element Element,
) Content {
	return createContentInternally(hash, triggersOn, element, nil)
}

func createContentWithFees(
	hash hash.Hash,
	triggersOn time.Time,
	fees []expenses.Expense,
) Content {
	return createContentInternally(hash, triggersOn, nil, fees)
}

func createContentWithElementAndFees(
	hash hash.Hash,
	triggersOn time.Time,
	element Element,
	fees []expenses.Expense,
) Content {
	return createContentInternally(hash, triggersOn, element, fees)
}

func createContentInternally(
	hash hash.Hash,
	triggersOn time.Time,
	element Element,
	fees []expenses.Expense,
) Content {
	out := content{
		hash:       hash,
		triggersOn: triggersOn,
		element:    element,
		fees:       fees,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// TriggersOn returns the triggersOn
func (obj *content) TriggersOn() time.Time {
	return obj.triggersOn
}

// HasElement returns true if there is an element, false otherwise
func (obj *content) HasElement() bool {
	return obj.element != nil
}

// Element returns the element
func (obj *content) Element() Element {
	return obj.element
}

// HasFees returns true if there fees, false otherwise
func (obj *content) HasFees() bool {
	return obj.fees != nil
}

// Fees returns the fees, if any
func (obj *content) Fees() []expenses.Expense {
	return obj.fees
}
