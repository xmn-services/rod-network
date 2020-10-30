package expenses

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
)

// JSONContent represents a json content
type JSONContent struct {
	Amount    uint64          `json:"amount"`
	From      *bills.JSONBill `json:"from"`
	Cancel    *locks.JSONLock `json:"cancel"`
	Remaining *locks.JSONLock `json:"remaining"`
	CreatedOn time.Time       `json:"created_on"`
}

func createJSONContentFromContent(content Content) *JSONContent {
	billsAdapter := bills.NewAdapter()
	from := billsAdapter.ToJSON(content.From())

	locksAdapter := locks.NewAdapter()
	cancel := locksAdapter.ToJSON(content.Cancel())
	var remaining *locks.JSONLock
	if content.HasRemaining() {
		remaining = locksAdapter.ToJSON(content.Remaining())
	}

	amount := content.Amount()
	createdOn := content.CreatedOn()
	return createJSONContent(amount, from, cancel, remaining, createdOn)
}

func createJSONContent(
	amount uint64,
	from *bills.JSONBill,
	cancel *locks.JSONLock,
	remaining *locks.JSONLock,
	createdOn time.Time,
) *JSONContent {
	out := JSONContent{
		Amount:    amount,
		From:      from,
		Cancel:    cancel,
		Remaining: remaining,
		CreatedOn: createdOn,
	}

	return &out
}
