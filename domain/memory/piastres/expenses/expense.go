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
	content    Content
	signatures [][]signature.RingSignature
}

func createExpenseFromJSON(ins *JSONExpense) (Expense, error) {
	content, err := createContentFromJSON(ins.Content)
	if err != nil {
		return nil, err
	}

	ringSigAdapter := signature.NewRingSignatureAdapter()
	signatures := [][]signature.RingSignature{}
	for _, oneSigList := range ins.Signatures {
		signaturesList := []signature.RingSignature{}
		for _, oneSigStr := range oneSigList {
			sig, err := ringSigAdapter.ToSignature(oneSigStr)
			if err != nil {
				return nil, err
			}

			signaturesList = append(signaturesList, sig)
		}

		signatures = append(signatures, signaturesList)
	}

	return NewBuilder().Create().WithContent(content).WithSignatures(signatures).Now()
}

func createExpense(
	immutable entities.Immutable,
	content Content,
	signatures [][]signature.RingSignature,
) Expense {
	out := expense{
		immutable:  immutable,
		content:    content,
		signatures: signatures,
	}

	return &out
}

// Hash returns the hash
func (obj *expense) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Content returns the content
func (obj *expense) Content() Content {
	return obj.content
}

// Signatures returns the signatures
func (obj *expense) Signatures() [][]signature.RingSignature {
	return obj.signatures
}

// CreatedOn returns the creation time
func (obj *expense) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *expense) MarshalJSON() ([]byte, error) {
	ins := createJSONExpenseFromExpense(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *expense) UnmarshalJSON(data []byte) error {
	ins := new(JSONExpense)
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
	obj.content = insExpense.content
	obj.signatures = insExpense.signatures
	return nil
}
