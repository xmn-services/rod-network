package transactions

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions/addresses"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type transaction struct {
	immutable entities.Immutable
	address   addresses.Address
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

	if js.Address != nil {
		addressAdapter := addresses.NewAdapter()
		address, err := addressAdapter.ToAddress(js.Address)
		if err != nil {
			return nil, err
		}

		builder.WithAddress(address)
	}

	return builder.Now()
}

func createTransactionWithAddress(
	immutable entities.Immutable,
	address addresses.Address,
) Transaction {
	return createTransactionInternally(immutable, address, nil, nil)
}

func createTransactionWithBucket(
	immutable entities.Immutable,
	bucket *hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, nil, bucket, nil)
}

func createTransactionWithFees(
	immutable entities.Immutable,
	fees []expenses.Expense,
) Transaction {
	return createTransactionInternally(immutable, nil, nil, fees)
}

func createTransactionWithAddressAndBucket(
	immutable entities.Immutable,
	address addresses.Address,
	bucket *hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, address, bucket, nil)
}

func createTransactionWithAddressAndFees(
	immutable entities.Immutable,
	address addresses.Address,
	fees []expenses.Expense,
) Transaction {
	return createTransactionInternally(immutable, address, nil, fees)
}

func createTransactionWithBucketAndFees(
	immutable entities.Immutable,
	bucket *hash.Hash,
	fees []expenses.Expense,
) Transaction {
	return createTransactionInternally(immutable, nil, bucket, fees)
}

func createTransactionWithAddressAndBucketAndFees(
	immutable entities.Immutable,
	address addresses.Address,
	bucket *hash.Hash,
	fees []expenses.Expense,
) Transaction {
	return createTransactionInternally(immutable, address, bucket, fees)
}

func createTransactionInternally(
	immutable entities.Immutable,
	address addresses.Address,
	bucket *hash.Hash,
	fees []expenses.Expense,
) Transaction {
	out := transaction{
		immutable: immutable,
		address:   address,
		bucket:    bucket,
		fees:      fees,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// HasAddress returns true if there is an address, false otherwise
func (obj *transaction) HasAddress() bool {
	return obj.address != nil
}

// Address returns the address, if any
func (obj *transaction) Address() addresses.Address {
	return obj.address
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

// MarshalJSON converts the instance to JSON
func (obj *transaction) MarshalJSON() ([]byte, error) {
	ins := createJSONTransactionFromTransaction(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *transaction) UnmarshalJSON(data []byte) error {
	ins := new(JSONTransaction)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createTransactionFromJSON(ins)
	if err != nil {
		return err
	}

	insTransaction := pr.(*transaction)
	obj.immutable = insTransaction.immutable
	obj.address = insTransaction.address
	obj.bucket = insTransaction.bucket
	obj.fees = insTransaction.fees
	return nil
}
