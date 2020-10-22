package mined

import (
	"reflect"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestBlock_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	blk, _ := hashAdapter.FromBytes([]byte("to build the block hash..."))
	mining := "334524352345"
	createdOn := time.Now().UTC()

	block, err := NewBuilder().Create().WithHash(*hsh).WithBlock(*blk).WithMining(mining).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !block.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !block.Block().Compare(*blk) {
		t.Errorf("the block hash is invalid")
		return
	}

	if !reflect.DeepEqual(block.Mining(), mining) {
		t.Errorf("the mining results are invalid")
		return
	}

	if !block.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), block.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(block)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBlock, err := repository.Retrieve(block.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, block, retBlock)
}
