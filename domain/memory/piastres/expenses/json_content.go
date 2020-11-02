package expenses

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
)

// JSONContent represents a json content
type JSONContent struct {
	From      []*bills.JSONBill `json:"from"`
	To        *bills.JSONBill   `json:"to"`
	Remaining *bills.JSONBill   `json:"remaining"`
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

	to := billsAdapter.ToJSON(content.To())

	var remaining *bills.JSONBill
	if content.HasRemaining() {
		remaining = billsAdapter.ToJSON(content.Remaining())
	}

	createdOn := content.CreatedOn()
	return createJSONContent(fromJS, to, remaining, createdOn)
}

func createJSONContent(
	from []*bills.JSONBill,
	to *bills.JSONBill,
	remaining *bills.JSONBill,
	createdOn time.Time,
) *JSONContent {
	out := JSONContent{
		From:      from,
		To:        to,
		Remaining: remaining,
		CreatedOn: createdOn,
	}

	return &out
}
