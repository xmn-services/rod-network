package blocks

import (
	"fmt"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

func TestBlock_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))

	data := [][]byte{}
	for i := 0; i < 5; i++ {
		str := fmt.Sprintf("to build the %d trx hash...", i)
		oneTrx, _ := hashAdapter.FromBytes([]byte(str))
		data = append(data, oneTrx.Bytes())
	}

	trx, _ := hashtree.NewBuilder().Create().WithBlocks(data).Now()
	amount := uint(len(data))
	createdOn := time.Now().UTC()

	block, err := NewBuilder().Create().WithHash(*hsh).WithTransactions(trx).WithAmount(amount).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !block.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !block.Transactions().Head().Compare(trx.Head()) {
		t.Errorf("the hashtree is invalid")
		return
	}

	if block.Amount() != amount {
		t.Errorf("the amount was expected to be %d, %d returned", amount, block.Amount())
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
