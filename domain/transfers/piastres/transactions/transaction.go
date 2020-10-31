package transactions

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type transaction struct {
	immutable entities.Immutable
	bucket    *hash.Hash
	fees      []hash.Hash
}

func createTransactionFromJSON(ins *jsonTransaction) (Transaction, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	builder := NewBuilder().
		Create().
		WithHash(*hsh).
		CreatedOn(ins.CreatedOn)

	if len(ins.Fees) > 0 {
		fees := []hash.Hash{}
		for _, oneFee := range ins.Fees {
			fee, err := hashAdapter.FromString(oneFee)
			if err != nil {
				return nil, err
			}

			fees = append(fees, *fee)
		}

		builder.WithFees(fees)
	}

	if ins.Bucket != "" {
		bucket, err := hashAdapter.FromString(ins.Bucket)
		if err != nil {
			return nil, err
		}

		builder.WithBucket(*bucket)
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
	fees []hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, nil, fees)
}

func createTransactionWithBucketAndFees(
	immutable entities.Immutable,
	bucket *hash.Hash,
	fees []hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, bucket, fees)
}

func createTransactionInternally(
	immutable entities.Immutable,
	bucket *hash.Hash,
	fees []hash.Hash,
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

// CreatedOn returns the creation time
func (obj *transaction) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasBucket returns true if the transaction is a bucket, false otherwise
func (obj *transaction) HasBucket() bool {
	return obj.bucket != nil
}

// Bucket returns the bucket hash, if any
func (obj *transaction) Bucket() *hash.Hash {
	return obj.bucket
}

// HasFees retruns true if there is fees, false otherwise
func (obj *transaction) HasFees() bool {
	return obj.fees != nil
}

// Fees returns the fees, if any
func (obj *transaction) Fees() []hash.Hash {
	return obj.fees
}

// MarshalJSON converts the instance to JSON
func (obj *transaction) MarshalJSON() ([]byte, error) {
	ins := createJSONTransactionFromTransaction(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *transaction) UnmarshalJSON(data []byte) error {
	ins := new(jsonTransaction)
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
	obj.fees = insTransaction.fees
	obj.bucket = insTransaction.bucket
	return nil
}
