package locks

import (
	"reflect"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

func TestLock_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))

	blocks := [][]byte{
		[]byte("to build the first shareholder hash..."),
		[]byte("to build the first shareholder hash..."),
	}

	holders, _ := hashtree.NewBuilder().Create().WithBlocks(blocks).Now()
	treeshold := uint(12)
	amount := uint(len(blocks))
	createdOn := time.Now().UTC()

	lock, err := NewBuilder().Create().WithHash(*hsh).WithShareHolders(holders).WithTreeshold(treeshold).WithAmount(amount).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !lock.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !reflect.DeepEqual(holders, lock.ShareHolders()) {
		t.Errorf("the shareholders are invalid")
		return
	}

	if lock.Treeshold() != treeshold {
		t.Errorf("the treeshold is invalid, expected: %d, returned: %d", treeshold, lock.Treeshold())
		return
	}

	if lock.Amount() != amount {
		t.Errorf("the amount is invalid, expected: %d, returned: %d", amount, lock.Amount())
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

func TestLock_withZeroAmount_returnsError(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))

	blocks := [][]byte{
		[]byte("to build the first shareholder hash..."),
		[]byte("to build the first shareholder hash..."),
	}

	holders, _ := hashtree.NewBuilder().Create().WithBlocks(blocks).Now()
	treeshold := uint(12)
	createdOn := time.Now().UTC()

	_, err := NewBuilder().Create().WithHash(*hsh).WithShareHolders(holders).WithTreeshold(treeshold).WithAmount(0).CreatedOn(createdOn).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestLock_withAmountTooBig_returnsError(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))

	blocks := [][]byte{
		[]byte("to build the first shareholder hash..."),
		[]byte("to build the first shareholder hash..."),
	}

	holders, _ := hashtree.NewBuilder().Create().WithBlocks(blocks).Now()
	treeshold := uint(12)
	amount := uint(len(blocks) + 1)
	createdOn := time.Now().UTC()

	_, err := NewBuilder().Create().WithHash(*hsh).WithShareHolders(holders).WithTreeshold(treeshold).WithAmount(amount).CreatedOn(createdOn).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
