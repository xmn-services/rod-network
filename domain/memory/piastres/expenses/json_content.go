package expenses

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
)

// JSONContent represents a json content
type JSONContent struct {
	Amount    uint64            `json:"amount"`
	From      []*bills.JSONBill `json:"from"`
	Remaining *locks.JSONLock   `json:"remaining"`
	CreatedOn time.Time         `json:"created_on"`
}

func createJSONContentFromContent(content Content) *JSONContent {
	from := content.From()
	fromJS := []*bills.JSONBill{}
	billsAdapter := bills.NewAdapter()
	for _, oneBill := range from {
		single := billsAdapter.ToJSON(oneBill)
		fromJS = append(fromJS, single)
	}

	locksAdapter := locks.NewAdapter()
	var remaining *locks.JSONLock
	if content.HasRemaining() {
		remaining = locksAdapter.ToJSON(content.Remaining())
	}

	amount := content.Amount()
	createdOn := content.CreatedOn()
	return createJSONContent(amount, fromJS, remaining, createdOn)
}

func createJSONContent(
	amount uint64,
	from []*bills.JSONBill,
	remaining *locks.JSONLock,
	createdOn time.Time,
) *JSONContent {
	out := JSONContent{
		Amount:    amount,
		From:      from,
		Remaining: remaining,
		CreatedOn: createdOn,
	}

	return &out
}
