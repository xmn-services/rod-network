package locks

import (
	"reflect"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestLock_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))

	firstPubKey, _ := hashAdapter.FromBytes([]byte("to build the first pubkey hash..."))
	secondPubKey, _ := hashAdapter.FromBytes([]byte("to build the second pubkey hash..."))

	pubkeys := []hash.Hash{
		*firstPubKey,
		*secondPubKey,
	}

	createdOn := time.Now().UTC()

	lock, err := NewBuilder().Create().WithHash(*hsh).WithPublicKeys(pubkeys).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !lock.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !reflect.DeepEqual(pubkeys, lock.PublicKeys()) {
		t.Errorf("the pubkeys are invalid")
		return
	}

	if !lock.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), lock.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(lock)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retLock, err := repository.Retrieve(lock.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, lock, retLock)
}
