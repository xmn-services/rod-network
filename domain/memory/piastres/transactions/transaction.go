package transactions

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type transaction struct {
	immutable entities.Immutable
	content   Content
	signature signature.RingSignature
}

func createTransactionFromJSON(ins *JSONTransaction) (Transaction, error) {
	content, err := createContentFromJSON(ins.Content)
	if err != nil {
		return nil, err
	}

	signatureAdapter := signature.NewRingSignatureAdapter()
	sig, err := signatureAdapter.ToSignature(ins.Signature)
	if err != nil {
		return nil, err
	}

	ringSize := uint(len(sig.Ring()) - 1)
	return NewBuilder(ringSize).Create().CreatedOn(ins.CreatedOn).WithContent(content).WithSignature(sig).Now()
}

func createTransaction(
	immutable entities.Immutable,
	content Content,
	signature signature.RingSignature,
) Transaction {
	out := transaction{
		immutable: immutable,
		content:   content,
		signature: signature,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Content returns the content
func (obj *transaction) Content() Content {
	return obj.content
}

// Signature returns the signature
func (obj *transaction) Signature() signature.RingSignature {
	return obj.signature
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

	insTrx := pr.(*transaction)
	obj.immutable = insTrx.immutable
	obj.content = insTrx.content
	obj.signature = insTrx.signature
	return nil
}
