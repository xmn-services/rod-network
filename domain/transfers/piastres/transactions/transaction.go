package transactions

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type transaction struct {
	immutable         entities.Immutable
	signature         signature.RingSignature
	triggersOn        time.Time
	executesOnTrigger bool
	expense           *hash.Hash
	cancel            *hash.Hash
	fees              *hash.Hash
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

	if ins.Fees != "" {
		fees, err := hashAdapter.FromString(ins.Fees)
		if err != nil {
			return nil, err
		}

		builder.WithFees(*fees)
	}

	if ins.Expense != "" {
		expense, err := hashAdapter.FromString(ins.Expense)
		if err != nil {
			return nil, err
		}

		builder.WithExpense(*expense)
	}

	if ins.Cancel != "" {
		cancel, err := hashAdapter.FromString(ins.Cancel)
		if err != nil {
			return nil, err
		}

		builder.WithCancel(*cancel)
	}

	if ins.ExecutesOnTrigger {
		builder.ExecutesOnTrigger()
	}

	return builder.Now()
}

func createTransactionWithExpense(
	immutable entities.Immutable,
	signature signature.RingSignature,
	triggersOn time.Time,
	executesOnTrigger bool,
	expense *hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, signature, triggersOn, executesOnTrigger, expense, nil, nil)
}

func createTransactionWithExpenseAndFees(
	immutable entities.Immutable,
	signature signature.RingSignature,
	triggersOn time.Time,
	executesOnTrigger bool,
	expense *hash.Hash,
	fees *hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, signature, triggersOn, executesOnTrigger, expense, nil, fees)
}

func createTransactionWithCancel(
	immutable entities.Immutable,
	signature signature.RingSignature,
	triggersOn time.Time,
	executesOnTrigger bool,
	cancel *hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, signature, triggersOn, executesOnTrigger, nil, cancel, nil)
}

func createTransactionWithCancelAndFees(
	immutable entities.Immutable,
	signature signature.RingSignature,
	triggersOn time.Time,
	executesOnTrigger bool,
	cancel *hash.Hash,
	fees *hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, signature, triggersOn, executesOnTrigger, nil, cancel, fees)
}

func createTransactionInternally(
	immutable entities.Immutable,
	signature signature.RingSignature,
	triggersOn time.Time,
	executesOnTrigger bool,
	expense *hash.Hash,
	cancel *hash.Hash,
	fees *hash.Hash,
) Transaction {
	out := transaction{
		immutable:         immutable,
		signature:         signature,
		triggersOn:        triggersOn,
		executesOnTrigger: executesOnTrigger,
		expense:           expense,
		cancel:            cancel,
		fees:              fees,
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

// ExecutesOnTrigger returns true if executes on trigger, false if cancel
func (obj *transaction) ExecutesOnTrigger() bool {
	return obj.executesOnTrigger
}

// CreatedOn returns the creation time
func (obj *transaction) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// IsExpense returns true if the transaction is an expense, false otherwise
func (obj *transaction) IsExpense() bool {
	return obj.expense != nil
}

// Expense returns the expense hash, if any
func (obj *transaction) Expense() *hash.Hash {
	return obj.expense
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
func (obj *transaction) Fees() *hash.Hash {
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
	obj.executesOnTrigger = insTransaction.executesOnTrigger
	obj.fees = insTransaction.fees
	obj.expense = insTransaction.expense
	obj.cancel = insTransaction.cancel
	return nil
}
