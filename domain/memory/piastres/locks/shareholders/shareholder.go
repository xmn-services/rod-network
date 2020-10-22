package shareholders

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type shareHolder struct {
	immutable entities.Immutable
	key       hash.Hash
	power     uint
}

func createShareHolderFromJSON(ins *JSONShareHolder) (ShareHolder, error) {
	hashAdapter := hash.NewAdapter()
	key, err := hashAdapter.FromString(ins.Key)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithKey(*key).
		WithPower(ins.Power).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createShareHolder(
	immutable entities.Immutable,
	key hash.Hash,
	power uint,
) ShareHolder {
	out := shareHolder{
		immutable: immutable,
		key:       key,
		power:     power,
	}

	return &out
}

// Hash returns the hash
func (obj *shareHolder) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Key returns the public key
func (obj *shareHolder) Key() hash.Hash {
	return obj.key
}

// Power returns the power
func (obj *shareHolder) Power() uint {
	return obj.power
}

// CreatedOn returns the creation time
func (obj *shareHolder) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *shareHolder) MarshalJSON() ([]byte, error) {
	ins := createJSONShareHolderFromShareHolder(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *shareHolder) UnmarshalJSON(data []byte) error {
	ins := new(JSONShareHolder)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createShareHolderFromJSON(ins)
	if err != nil {
		return err
	}

	insShareHolder := pr.(*shareHolder)
	obj.immutable = insShareHolder.immutable
	obj.key = insShareHolder.key
	obj.power = insShareHolder.power
	return nil
}
