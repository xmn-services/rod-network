package transactions

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/cancels"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
)

// JSONContent represents a json content
type JSONContent struct {
	TriggersOn time.Time               `json:"triggers_on"`
	Fees       []*expenses.JSONExpense `json:"fees"`
	Bucket     string                  `json:"bucket"`
	Cancel     *cancels.JSONCancel     `json:"cancel"`
}

func createJSONContentFromContent(content Content) *JSONContent {
	triggersOn := content.TriggersOn()
	expenseAdapter := expenses.NewAdapter()
	jsonFees := []*expenses.JSONExpense{}
	if content.HasFees() {
		fees := content.Fees()
		for _, oneFee := range fees {
			fee := expenseAdapter.ToJSON(oneFee)
			jsonFees = append(jsonFees, fee)
		}
	}

	cancelAdapter := cancels.NewAdapter()
	var jsonCancel *cancels.JSONCancel
	bucket := ""
	if content.HasElement() {
		element := content.Element()
		if element.IsBucket() {
			bucket = element.Bucket().String()
		}

		if element.IsCancel() {
			cancel := element.Cancel()
			jsonCancel = cancelAdapter.ToJSON(cancel)
		}
	}

	return createJSONContent(triggersOn, jsonFees, bucket, jsonCancel)
}

func createJSONContent(
	triggersOn time.Time,
	fees []*expenses.JSONExpense,
	bucket string,
	cancel *cancels.JSONCancel,
) *JSONContent {
	out := JSONContent{
		TriggersOn: triggersOn,
		Fees:       fees,
		Bucket:     bucket,
		Cancel:     cancel,
	}

	return &out
}
