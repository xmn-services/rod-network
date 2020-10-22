package transactions

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/cancels"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
)

// JSONContent represents a json content
type JSONContent struct {
	TriggersOn        time.Time             `json:"triggers_on"`
	ExecutesOnTrigger bool                  `json:"executes_on_trigger"`
	Fees              *expenses.JSONExpense `json:"fees"`
	Expense           *expenses.JSONExpense `json:"expense"`
	Cancel            *cancels.JSONCancel   `json:"cancel"`
}

func createJSONContentFromContent(content Content) *JSONContent {
	triggersOn := content.TriggersOn()
	executesOnTrigger := content.ExecutesOnTrigger()

	expenseAdapter := expenses.NewAdapter()
	var fees *expenses.JSONExpense
	if content.HasFees() {
		fees = expenseAdapter.ToJSON(content.Fees())
	}

	cancelAdapter := cancels.NewAdapter()
	var cancel *cancels.JSONCancel
	if content.IsCancel() {
		cancel = cancelAdapter.ToJSON(content.Cancel())
	}

	var expense *expenses.JSONExpense
	if content.IsExpense() {
		expense = expenseAdapter.ToJSON(content.Expense())
	}

	return createJSONContent(triggersOn, executesOnTrigger, fees, expense, cancel)
}

func createJSONContent(
	triggersOn time.Time,
	executesOnTrigger bool,
	fees *expenses.JSONExpense,
	expense *expenses.JSONExpense,
	cancel *cancels.JSONCancel,
) *JSONContent {
	out := JSONContent{
		TriggersOn:        triggersOn,
		ExecutesOnTrigger: executesOnTrigger,
		Fees:              fees,
		Expense:           expense,
		Cancel:            cancel,
	}

	return &out
}
