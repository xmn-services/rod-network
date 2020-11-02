package locks

import (
	"encoding/json"
	"testing"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestLock_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pkFactory := signature.NewPrivateKeyFactory()
	firstPK := pkFactory.Create()
	secondPK := pkFactory.Create()
	thirdPK := pkFactory.Create()

	firstHash, _ := hashAdapter.FromBytes([]byte(firstPK.PublicKey().String()))
	secondHash, _ := hashAdapter.FromBytes([]byte(secondPK.PublicKey().String()))
	thirdHash, _ := hashAdapter.FromBytes([]byte(thirdPK.PublicKey().String()))

	pubKeys := []hash.Hash{
		*firstHash,
		*secondHash,
		*thirdHash,
	}

	lockIns := CreateLockForTests(pubKeys)

	if len(lockIns.PublicKeys()) != len(pubKeys) {
		t.Errorf("%d PublicKey were expected, %d returned", len(pubKeys), len(lockIns.PublicKeys()))
		return
	}

	ringPubKeys := []signature.PublicKey{
		firstPK.PublicKey(),
		secondPK.PublicKey(),
		thirdPK.PublicKey(),
	}

	sig, err := firstPK.RingSign(lockIns.Hash().String(), ringPubKeys)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = lockIns.Unlock(sig)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// repository and service:
	repository, service := CreateRepositoryServiceForTests()
	err = service.Save(lockIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retLock, err := repository.Retrieve(lockIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !lockIns.Hash().Compare(retLock.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(lockIns)
	if err != nil {
		t.Errorf("the Lock instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(lock)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a Lock instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, ins, lockIns)
}
