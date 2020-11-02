package transactions

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions/addresses"
)

// JSONTransaction represents a json transaction
type JSONTransaction struct {
	Address   *addresses.JSONAddress  `json:"address"`
	Fees      []*expenses.JSONExpense `json:"fees"`
	Bucket    string                  `json:"bucket"`
	CreatedOn time.Time               `json:"created_on"`
}

func createJSONTransactionFromTransaction(transaction Transaction) *JSONTransaction {
	var address *addresses.JSONAddress
	if transaction.HasAddress() {
		addressAdapter := addresses.NewAdapter()
		addr := transaction.Address()
		address = addressAdapter.ToJSON(addr)
	}

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
	return createJSONTransaction(address, jsonFees, bucket, createdOn)
}

func createJSONTransaction(
	address *addresses.JSONAddress,
	fees []*expenses.JSONExpense,
	bucket string,
	createdOn time.Time,
) *JSONTransaction {
	out := JSONTransaction{
		Address:   address,
		Fees:      fees,
		Bucket:    bucket,
		CreatedOn: createdOn,
	}

	return &out
}
