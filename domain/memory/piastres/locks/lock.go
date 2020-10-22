package locks

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
)

type lock struct {
	immutable    entities.Immutable
	shareholders []shareholders.ShareHolder
	mp           map[string]shareholders.ShareHolder
	treshold     uint
}

func createLockFromJSON(ins *JSONLock) (Lock, error) {
	holders := []shareholders.ShareHolder{}
	shareHolderAdapter := shareholders.NewAdapter()
	for _, oneJSONHolder := range ins.ShareHolders {
		holder, err := shareHolderAdapter.ToShareHolder(oneJSONHolder)
		if err != nil {
			return nil, err
		}

		holders = append(holders, holder)
	}

	return NewBuilder().
		Create().
		WithShareHolders(holders).
		WithTreeshold(ins.Treeshold).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createLock(
	immutable entities.Immutable,
	shareholders []shareholders.ShareHolder,
	mp map[string]shareholders.ShareHolder,
	treshold uint,
) Lock {
	out := lock{
		immutable:    immutable,
		shareholders: shareholders,
		mp:           mp,
		treshold:     treshold,
	}

	return &out
}

// Hash returns the hash
func (obj *lock) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// ShareHolders returns the shareholders
func (obj *lock) ShareHolders() []shareholders.ShareHolder {
	return obj.shareholders
}

// Treeshold returns the treshold
func (obj *lock) Treeshold() uint {
	return obj.treshold
}

// CreatedOn returns the creation time
func (obj *lock) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// Unlock verifies if the given combination unlocks the current lock
func (obj *lock) Unlock(signatures []signature.RingSignature) error {
	hashAdapter := hash.NewAdapter()
	getShareHolder := func(ring []signature.PublicKey, shareholders map[string]shareholders.ShareHolder) (shareholders.ShareHolder, error) {
		for _, oneRingKey := range ring {
			ringKeyHash, err := hashAdapter.FromBytes([]byte(oneRingKey.String()))
			if err != nil {
				return nil, err
			}

			ringKeyStr := ringKeyHash.String()
			if shareHolder, ok := shareholders[ringKeyStr]; ok {
				return shareHolder, nil
			}
		}

		return nil, errors.New("at least 1 Signature that does not fit our []ShareHolder was provided")
	}

	power := uint(0)
	already := map[string]shareholders.ShareHolder{}
	for _, oneSignature := range signatures {
		ring := oneSignature.Ring()
		shareHolder, err := getShareHolder(ring, obj.mp)
		if err != nil {
			return err
		}

		ringKey := shareHolder.Key().String()
		if _, ok := already[ringKey]; ok {
			continue
		}

		// validate the signature:
		if !oneSignature.Verify(obj.immutable.Hash().String()) {
			return errors.New("at least 1 Signature was invalid")
		}

		// add the power:
		power += shareHolder.Power()
		already[ringKey] = shareHolder
	}

	if power < obj.Treeshold() {
		str := fmt.Sprintf("the Lock could not be unlocked because it doesn't have enough power (provided: %d, expected: %d)", power, obj.Treeshold())
		return errors.New(str)
	}

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
	obj.shareholders = insLock.shareholders
	obj.treshold = insLock.treshold
	obj.mp = insLock.mp
	return nil
}
