package locks

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type lock struct {
	immutable entities.Immutable
	pubKeys   []hash.Hash
	mpKeys    map[string]hash.Hash
}

func createLockFromJSON(ins *JSONLock) (Lock, error) {
	hashAdapter := hash.NewAdapter()
	publicKeys := []hash.Hash{}
	for _, onePubKeyStr := range ins.PublicKeys {
		hsh, err := hashAdapter.FromString(onePubKeyStr)
		if err != nil {
			return nil, err
		}

		publicKeys = append(publicKeys, *hsh)
	}

	return NewBuilder().
		Create().
		WithPublicKeys(publicKeys).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createLock(
	immutable entities.Immutable,
	pubKeys []hash.Hash,
	mpKeys map[string]hash.Hash,
) Lock {
	out := lock{
		immutable: immutable,
		pubKeys:   pubKeys,
		mpKeys:    mpKeys,
	}

	return &out
}

// Hash returns the hash
func (obj *lock) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// PublicKeys returns the publickeys
func (obj *lock) PublicKeys() []hash.Hash {
	return obj.pubKeys
}

// CreatedOn returns the creation time
func (obj *lock) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// Unlock verifies if the given combination unlocks the current lock
func (obj *lock) Unlock(signature signature.RingSignature) error {
	ring := signature.Ring()
	ringMp := map[string]hash.Hash{}
	hashAdapter := hash.NewAdapter()
	for _, oneRingPubKey := range ring {
		ringHash, err := hashAdapter.FromBytes([]byte(oneRingPubKey.String()))
		if err != nil {
			return err
		}

		keyname := ringHash.String()
		if _, ok := obj.mpKeys[keyname]; !ok {
			return errors.New("at least 1 PublicKey contained in the RingSignature was not in the lock's PublicKey list")
		}

		ringMp[keyname] = *ringHash
	}

	if len(ringMp) != len(obj.mpKeys) {
		str := fmt.Sprintf("the RingSignature provided contains %d unique PublicKey, while the Lock contains %d unique PublicKey", len(ring), len(obj.mpKeys))
		return errors.New(str)
	}

	// validate the signature:
	if !signature.Verify(obj.immutable.Hash().String()) {
		return errors.New("the signature is invalid")
	}

	// the signature is valid:
	return nil
}

// MarshalJSON converts the instance to JSON
func (obj *lock) MarshalJSON() ([]byte, error) {
	ins := createJSONLockFromLock(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *lock) UnmarshalJSON(data []byte) error {
	ins := new(JSONLock)
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
	obj.mpKeys = insLock.mpKeys
	return nil
}
