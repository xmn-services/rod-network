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
	from       []hash.Hash
	to         hash.Hash
	signatures []signature.RingSignature
	remaining  *hash.Hash
}

func createExpenseFromJSON(ins *jsonExpense) (Expense, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	from := []hash.Hash{}
	for _, oneStr := range ins.From {
		oneFrom, err := hashAdapter.FromString(oneStr)
		if err != nil {
			return nil, err
		}

		from = append(from, *oneFrom)
	}

	to, err := hashAdapter.FromString(ins.To)
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
		From(from).
		To(*to).
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
	from []hash.Hash,
	to hash.Hash,
	signatures []signature.RingSignature,
) Expense {
	return createExpenseInternally(immutable, from, to, signatures, nil)
}

func createExpenseWithRemaining(
	immutable entities.Immutable,
	from []hash.Hash,
	to hash.Hash,
	signatures []signature.RingSignature,
	remaining *hash.Hash,
) Expense {
	return createExpenseInternally(immutable, from, to, signatures, remaining)
}

func createExpenseInternally(
	immutable entities.Immutable,
	from []hash.Hash,
	to hash.Hash,
	signatures []signature.RingSignature,
	remaining *hash.Hash,
) Expense {
	out := expense{
		immutable:  immutable,
		from:       from,
		to:         to,
		signatures: signatures,
		remaining:  remaining,
	}

	return &out
}

// Hash returns the hash
func (obj *expense) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// From returns the from hash
func (obj *expense) From() []hash.Hash {
	return obj.from
}

// To returns the to hash
func (obj *expense) To() hash.Hash {
	return obj.to
}

// CreatedOn returns the creation time
func (obj *expense) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
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
	obj.from = insExpense.from
	obj.to = insExpense.to
	obj.signatures = insExpense.signatures
	obj.remaining = insExpense.remaining
	return nil
}
