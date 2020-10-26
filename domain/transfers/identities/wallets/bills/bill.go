package bills

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type bill struct {
	immutable entities.Immutable
	bill      hash.Hash
	pks       []signature.PrivateKey
}

func createBillFromJSON(ins *jsonBill) (Bill, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	bill, err := hashAdapter.FromString(ins.Bill)
	if err != nil {
		return nil, err
	}

	pks := []signature.PrivateKey{}
	pkAdapter := signature.NewPrivateKeyAdapter()
	for _, oneStr := range ins.PrivateKeys {
		pk, err := pkAdapter.ToPrivateKey(oneStr)
		if err != nil {
			return nil, err
		}

		pks = append(pks, pk)
	}

	return NewBuilder().Create().
		WithHash(*hsh).
		WithBill(*bill).
		WithPrivateKeys(pks).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createBill(
	immutable entities.Immutable,
	bll hash.Hash,
	pks []signature.PrivateKey,
) Bill {
	out := bill{
		immutable: immutable,
		bill:      bll,
		pks:       pks,
	}

	return &out
}

// Hash returns the hash
func (obj *bill) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Bill returns the bill
func (obj *bill) Bill() hash.Hash {
	return obj.bill
}

// PrivateKeys returns the privateKeys
func (obj *bill) PrivateKeys() []signature.PrivateKey {
	return obj.pks
}

// CreatedOn returns the creation time
func (obj *bill) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *bill) MarshalJSON() ([]byte, error) {
	ins := createJSONBillFromBill(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *bill) UnmarshalJSON(data []byte) error {
	ins := new(jsonBill)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createBillFromJSON(ins)
	if err != nil {
		return err
	}

	insBill := pr.(*bill)
	obj.immutable = insBill.immutable
	obj.bill = insBill.bill
	obj.pks = insBill.pks
	return nil
}
