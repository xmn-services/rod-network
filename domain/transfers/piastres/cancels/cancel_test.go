package cancels

import (
	"reflect"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestCancel_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	expense, _ := hashAdapter.FromBytes([]byte("to build the expense hash..."))
	lock, _ := hashAdapter.FromBytes([]byte("to build the lock hash..."))
	createdOn := time.Now().UTC()

	pk := signature.NewPrivateKeyFactory().Create()
	ring := []signature.PublicKey{
		pk.PublicKey(),
	}

	sig, _ := pk.RingSign(hsh.String(), ring)
	signatures := []signature.RingSignature{
		sig,
	}

	cancel, err := NewBuilder().Create().WithHash(*hsh).WithExpense(*expense).WithLock(*lock).WithSignatures(signatures).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !cancel.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !cancel.Expense().Compare(*expense) {
		t.Errorf("the expense hash is invalid")
		return
	}

	if !cancel.Lock().Compare(*lock) {
		t.Errorf("the lock hash is invalid")
		return
	}

	if !reflect.DeepEqual(signatures, cancel.Signatures()) {
		t.Errorf("the ring signatures are invalid")
		return
	}

	if !cancel.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), cancel.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(cancel)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retPeer, err := repository.Retrieve(cancel.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, cancel, retPeer)
}
