package transactions

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
)

// JSONTransaction represents a json transaction
type JSONTransaction struct {
	Fees      []*expenses.JSONExpense `json:"fees"`
	Bucket    string                  `json:"bucket"`
	CreatedOn time.Time               `json:"created_on"`
}

func createJSONTransactionFromTransaction(transaction Transaction) *JSONTransaction {
	expenseAdapter := expenses.NewAdapter()
	jsonFees := []*expenses.JSONExpense{}
	if transaction.HasFees() {
		fees := transaction.Fees()
		for _, oneFee := range fees {
			fee := expenseAdapter.ToJSON(oneFee)
			jsonFees = append(jsonFees, fee)
		}
	}

	bucket := ""
	if transaction.HasBucket() {
		bucket = transaction.Bucket().String()
	}

	createdOn := transaction.CreatedOn()
	return createJSONTransaction(jsonFees, bucket, createdOn)
}

func createJSONTransaction(
	fees []*expenses.JSONExpense,
	bucket string,
	createdOn time.Time,
) *JSONTransaction {
	out := JSONTransaction{
		Fees:      fees,
		Bucket:    bucket,
		CreatedOn: createdOn,
	}

	return &out
}
