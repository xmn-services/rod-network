package locks

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type lock struct {
	immutable entities.Immutable
	pubKeys   []hash.Hash
}

func createLockFromJSON(ins *jsonLock) (Lock, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	pubKeys := []hash.Hash{}
	for _, onePubKeyStr := range ins.PublicKeys {
		pubKey, err := hashAdapter.FromString(onePubKeyStr)
		if err != nil {
			return nil, err
		}

		pubKeys = append(pubKeys, *pubKey)
	}

	return NewBuilder().
		Create().
		WithHash(*hsh).
		WithPublicKeys(pubKeys).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createLock(
	immutable entities.Immutable,
	pubKeys []hash.Hash,
) Lock {
	out := lock{
		immutable: immutable,
		pubKeys:   pubKeys,
	}

	return &out
}

// Hash returns the hash
func (obj *lock) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// PublicKeys returns the public keys
func (obj *lock) PublicKeys() []hash.Hash {
	return obj.pubKeys
}

// CreatedOn returns the creation time
func (obj *lock) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *lock) MarshalJSON() ([]byte, error) {
	ins := createJSONLockFromLock(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *lock) UnmarshalJSON(data []byte) error {
	ins := new(jsonLock)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createLockFromJSON(ins)
	if err != nil {
		return err
	}

	insLock := pr.(*lock)
	obj.immutable = insLock.immutable
	obj.pubKeys = insLock.pubKeys
	return nil
}
