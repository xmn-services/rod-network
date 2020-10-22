package cancels

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
)

type cancel struct {
	immutable  entities.Immutable
	expense    expenses.Expense
	lock       locks.Lock
	signatures []signature.RingSignature
}

func createCancelFromJSON(ins *JSONCancel) (Cancel, error) {
	expenseAdapter := expenses.NewAdapter()
	expense, err := expenseAdapter.ToExpense(ins.Expense)
	if err != nil {
		return nil, err
	}

	lockAdapter := locks.NewAdapter()
	lock, err := lockAdapter.ToLock(ins.Lock)
	if err != nil {
		return nil, err
	}

	ringSigAdapter := signature.NewRingSignatureAdapter()
	signatures := []signature.RingSignature{}
	for _, oneSigStr := range ins.Signatures {
		sig, err := ringSigAdapter.ToSignature(oneSigStr)
		if err != nil {
			return nil, err
		}

		signatures = append(signatures, sig)
	}

	return NewBuilder().
		Create().
		WithExpense(expense).
		WithLock(lock).
		WithSignatures(signatures).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createCancel(
	immutable entities.Immutable,
	expense expenses.Expense,
	lock locks.Lock,
	signatures []signature.RingSignature,
) Cancel {
	out := cancel{
		immutable:  immutable,
		expense:    expense,
		lock:       lock,
		signatures: signatures,
	}

	return &out
}

// Hash returns the hash
func (obj *cancel) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Expense returns the expense
func (obj *cancel) Expense() expenses.Expense {
	return obj.expense
}

// Lock returns the lock
func (obj *cancel) Lock() locks.Lock {
	return obj.lock
}

// Signatures returns the signatures
func (obj *cancel) Signatures() []signature.RingSignature {
	return obj.signatures
}

// CreatedOn returns the creation time
func (obj *cancel) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *cancel) MarshalJSON() ([]byte, error) {
	ins := createJSONCancelFromCancel(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *cancel) UnmarshalJSON(data []byte) error {
	ins := new(JSONCancel)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createCancelFromJSON(ins)
	if err != nil {
		return err
	}

	insCancel := pr.(*cancel)
	obj.immutable = insCancel.immutable
	obj.expense = insCancel.expense
	obj.lock = insCancel.lock
	obj.signatures = insCancel.signatures
	return nil
}
