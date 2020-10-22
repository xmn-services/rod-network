package transactions

import (
	"time"
)

type jsonTransaction struct {
	Hash              string    `json:"hash"`
	Signature         string    `json:"signature"`
	TriggersOn        time.Time `json:"triggers_on"`
	ExecutesOnTrigger bool      `json:"executes_on_trigger"`
	Fees              string    `json:"fees"`
	Expense           string    `json:"expense"`
	Cancel            string    `json:"cancel"`
	CreatedOn         time.Time `json:"created_on"`
}

func createJSONTransactionFromTransaction(ins Transaction) *jsonTransaction {
	hash := ins.Hash().String()
	signature := ins.Signature().String()
	triggersOn := ins.TriggersOn()
	executesOnTrigger := ins.ExecutesOnTrigger()

	fees := ""
	if ins.HasFees() {
		fees = ins.Fees().String()
	}

	expense := ""
	if ins.IsExpense() {
		expense = ins.Expense().String()
	}

	cancel := ""
	if ins.IsCancel() {
		cancel = ins.Cancel().String()
	}

	createdOn := ins.CreatedOn()
	return createJSONTransaction(hash, signature, triggersOn, executesOnTrigger, fees, expense, cancel, createdOn)
}

func createJSONTransaction(
	hash string,
	signature string,
	triggersOn time.Time,
	executesOnTrigger bool,
	fees string,
	expense string,
	cancel string,
	createdOn time.Time,
) *jsonTransaction {
	out := jsonTransaction{
		Hash:              hash,
		Signature:         signature,
		TriggersOn:        triggersOn,
		ExecutesOnTrigger: executesOnTrigger,
		Fees:              fees,
		Expense:           expense,
		Cancel:            cancel,
		CreatedOn:         createdOn,
	}

	return &out
}
