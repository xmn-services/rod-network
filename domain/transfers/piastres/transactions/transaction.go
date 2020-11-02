package transactions

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type transaction struct {
	immutable entities.Immutable
	address   *hash.Hash
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

	if ins.Address != "" {
		address, err := hashAdapter.FromString(ins.Address)
		if err != nil {
			return nil, err
		}

		builder.WithAddress(*address)
	}

	return builder.Now()
}

func createTransactionWithAddress(
	immutable entities.Immutable,
	address *hash.Hash,
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
	fees []hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, nil, nil, fees)
}

func createTransactionWithAddressAndFees(
	immutable entities.Immutable,
	address *hash.Hash,
	fees []hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, address, nil, fees)
}

func createTransactionWithAddressAndBucket(
	immutable entities.Immutable,
	address *hash.Hash,
	bucket *hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, address, bucket, nil)
}

func createTransactionWithBucketAndFees(
	immutable entities.Immutable,
	bucket *hash.Hash,
	fees []hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, nil, bucket, fees)
}

func createTransactionWithAddressAndBucketAndFees(
	immutable entities.Immutable,
	address *hash.Hash,
	bucket *hash.Hash,
	fees []hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, address, bucket, fees)
}

func createTransactionInternally(
	immutable entities.Immutable,
	address *hash.Hash,
	bucket *hash.Hash,
	fees []hash.Hash,
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

// CreatedOn returns the creation time
func (obj *transaction) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasAddress returns true if the transaction contains an address, false otherwise
func (obj *transaction) HasAddress() bool {
	return obj.address != nil
}

// Address returns the address hash, if any
func (obj *transaction) Address() *hash.Hash {
	return obj.address
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
	obj.address = insTransaction.address
	obj.fees = insTransaction.fees
	obj.bucket = insTransaction.bucket
	return nil
}
