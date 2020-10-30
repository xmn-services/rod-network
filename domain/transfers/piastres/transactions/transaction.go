package transactions

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type transaction struct {
	immutable  entities.Immutable
	signature  signature.RingSignature
	triggersOn time.Time
	bucket     *hash.Hash
	cancel     *hash.Hash
	fees       []hash.Hash
}

func createTransactionFromJSON(ins *jsonTransaction) (Transaction, error) {
	hashAdapter := hash.NewAdapter()
	signatureAdapter := signature.NewRingSignatureAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	sig, err := signatureAdapter.ToSignature(ins.Signature)
	if err != nil {
		return nil, err
	}

	builder := NewBuilder().
		Create().
		WithHash(*hsh).
		TriggersOn(ins.TriggersOn).
		WithSignature(sig).
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

	if ins.Cancel != "" {
		cancel, err := hashAdapter.FromString(ins.Cancel)
		if err != nil {
			return nil, err
		}

		builder.WithCancel(*cancel)
	}

	return builder.Now()
}

func createTransactionWithBucket(
	immutable entities.Immutable,
	signature signature.RingSignature,
	triggersOn time.Time,
	bucket *hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, signature, triggersOn, bucket, nil, nil)
}

func createTransactionWithBucketAndFees(
	immutable entities.Immutable,
	signature signature.RingSignature,
	triggersOn time.Time,
	bucket *hash.Hash,
	fees []hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, signature, triggersOn, bucket, nil, fees)
}

func createTransactionWithCancel(
	immutable entities.Immutable,
	signature signature.RingSignature,
	triggersOn time.Time,
	cancel *hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, signature, triggersOn, nil, cancel, nil)
}

func createTransactionWithCancelAndFees(
	immutable entities.Immutable,
	signature signature.RingSignature,
	triggersOn time.Time,
	cancel *hash.Hash,
	fees []hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, signature, triggersOn, nil, cancel, fees)
}

func createTransactionWithFees(
	immutable entities.Immutable,
	signature signature.RingSignature,
	triggersOn time.Time,
	fees []hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, signature, triggersOn, nil, nil, fees)
}

func createTransactionInternally(
	immutable entities.Immutable,
	signature signature.RingSignature,
	triggersOn time.Time,
	bucket *hash.Hash,
	cancel *hash.Hash,
	fees []hash.Hash,
) Transaction {
	out := transaction{
		immutable:  immutable,
		signature:  signature,
		triggersOn: triggersOn,
		bucket:     bucket,
		cancel:     cancel,
		fees:       fees,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Signature returns the signature
func (obj *transaction) Signature() signature.RingSignature {
	return obj.signature
}

// TriggersOn returns the triggersOn time
func (obj *transaction) TriggersOn() time.Time {
	return obj.triggersOn
}

// CreatedOn returns the creation time
func (obj *transaction) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// IsBucket returns true if the transaction is a bucket, false otherwise
func (obj *transaction) IsBucket() bool {
	return obj.bucket != nil
}

// Bucket returns the bucket hash, if any
func (obj *transaction) Bucket() *hash.Hash {
	return obj.bucket
}

// IsCancel returns true if the transaction is a cancel, false otherwise
func (obj *transaction) IsCancel() bool {
	return obj.cancel != nil
}

// Cancel returns the cancel hash, if any
func (obj *transaction) Cancel() *hash.Hash {
	return obj.cancel
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
	obj.signature = insTransaction.signature
	obj.triggersOn = insTransaction.triggersOn
	obj.fees = insTransaction.fees
	obj.bucket = insTransaction.bucket
	obj.cancel = insTransaction.cancel
	return nil
}
