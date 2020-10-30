package expenses

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type expense struct {
	immutable  entities.Immutable
	amount     uint64
	from       hash.Hash
	cancel     hash.Hash
	signatures []signature.RingSignature
	remaining  *hash.Hash
}

func createExpenseFromJSON(ins *jsonExpense) (Expense, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	from, err := hashAdapter.FromString(ins.From)
	if err != nil {
		return nil, err
	}

	cancel, err := hashAdapter.FromString(ins.Cancel)
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

	builder := NewBuilder().Create().
		WithHash(*hsh).
		WithAmount(ins.Amount).
		From(*from).
		WithCancel(*cancel).
		WithSignatures(signatures).
		CreatedOn(ins.CreatedOn)

	if ins.Remaining != "" {
		remaining, err := hashAdapter.FromString(ins.Remaining)
		if err != nil {
			return nil, err
		}

		builder.WithRemaining(*remaining)
	}

	return builder.Now()
}

func createExpense(
	immutable entities.Immutable,
	amount uint64,
	from hash.Hash,
	cancel hash.Hash,
	signatures []signature.RingSignature,
) Expense {
	return createExpenseInternally(immutable, amount, from, cancel, signatures, nil)
}

func createExpenseWithRemaining(
	immutable entities.Immutable,
	amount uint64,
	from hash.Hash,
	cancel hash.Hash,
	signatures []signature.RingSignature,
	remaining *hash.Hash,
) Expense {
	return createExpenseInternally(immutable, amount, from, cancel, signatures, remaining)
}

func createExpenseInternally(
	immutable entities.Immutable,
	amount uint64,
	from hash.Hash,
	cancel hash.Hash,
	signatures []signature.RingSignature,
	remaining *hash.Hash,
) Expense {
	out := expense{
		immutable:  immutable,
		amount:     amount,
		from:       from,
		cancel:     cancel,
		signatures: signatures,
		remaining:  remaining,
	}

	return &out
}

// Hash returns the hash
func (obj *expense) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Amount returns the amount
func (obj *expense) Amount() uint64 {
	return obj.amount
}

// From returns the from hash
func (obj *expense) From() hash.Hash {
	return obj.from
}

// CreatedOn returns the creation time
func (obj *expense) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// Cancel returns the cancel hash
func (obj *expense) Cancel() hash.Hash {
	return obj.cancel
}

// Signatures returns the signatures
func (obj *expense) Signatures() []signature.RingSignature {
	return obj.signatures
}

// HasRemaining returns true if there is a remaining hash, false otherwise
func (obj *expense) HasRemaining() bool {
	return obj.remaining != nil
}

// Remaining returns the remaining hash, if any
func (obj *expense) Remaining() *hash.Hash {
	return obj.remaining
}

// MarshalJSON converts the instance to JSON
func (obj *expense) MarshalJSON() ([]byte, error) {
	ins := createJSONExpenseFromExpense(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *expense) UnmarshalJSON(data []byte) error {
	ins := new(jsonExpense)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createExpenseFromJSON(ins)
	if err != nil {
		return err
	}

	insExpense := pr.(*expense)
	obj.immutable = insExpense.immutable
	obj.amount = insExpense.amount
	obj.from = insExpense.from
	obj.cancel = insExpense.cancel
	obj.signatures = insExpense.signatures
	obj.remaining = insExpense.remaining
	return nil
}
