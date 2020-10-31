package transactions

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type transaction struct {
	immutable entities.Immutable
	bucket    *hash.Hash
	fees      []expenses.Expense
}

func createTransactionFromJSON(js *JSONTransaction) (Transaction, error) {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder().Create().CreatedOn(js.CreatedOn)
	if js.Bucket != "" {
		bucket, err := hashAdapter.FromString(js.Bucket)
		if err != nil {
			return nil, err
		}

		builder.WithBucket(*bucket)
	}

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

	return builder.Now()
}

func createTransactionWithBucket(
	immutable entities.Immutable,
	bucket *hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, bucket, nil)
}

func createTransactionWithFees(
	immutable entities.Immutable,
	fees []expenses.Expense,
) Transaction {
	return createTransactionInternally(immutable, nil, fees)
}

func createTransactionWithBucketAndFees(
	immutable entities.Immutable,
	bucket *hash.Hash,
	fees []expenses.Expense,
) Transaction {
	return createTransactionInternally(immutable, bucket, fees)
}

func createTransactionInternally(
	immutable entities.Immutable,
	bucket *hash.Hash,
	fees []expenses.Expense,
) Transaction {
	out := transaction{
		immutable: immutable,
		bucket:    bucket,
		fees:      fees,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// HasBucket returns true if there is a bucket, false otherwise
func (obj *transaction) HasBucket() bool {
	return obj.bucket != nil
}

// Bucket returns the bucket, if any
func (obj *transaction) Bucket() *hash.Hash {
	return obj.bucket
}

// HasFees returns true if there fees, false otherwise
func (obj *transaction) HasFees() bool {
	return obj.fees != nil
}

// Fees returns the fees, if any
func (obj *transaction) Fees() []expenses.Expense {
	return obj.fees
}

// CreatedOn returns the creation time
func (obj *transaction) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
