package states

import (
	"fmt"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

func TestState_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	chain, _ := hashAdapter.FromBytes([]byte("to build the chain hash..."))
	prev, _ := hashAdapter.FromBytes([]byte("to build the previous hash..."))

	height := uint(56)

	data := [][]byte{}
	for i := 0; i < 5; i++ {
		str := fmt.Sprintf("to build the %d hash...", i)
		data = append(data, []byte(str))
	}

	trx, _ := hashtree.NewBuilder().Create().WithBlocks(data).Now()
	createdOn := time.Now().UTC()

	amount := uint(len(data))

	state, err := NewBuilder().Create().WithHash(*hsh).WithChain(*chain).WithPrevious(*prev).WithHeight(height).WithTransactions(trx).WithAmount(amount).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !state.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !state.Chain().Compare(*chain) {
		t.Errorf("the chain hash is invalid")
		return
	}

	if !state.Previous().Compare(*prev) {
		t.Errorf("the previous hash is invalid")
		return
	}

	if state.Height() != height {
		t.Errorf("the height is invalid, expected: %d, returned: %d", height, state.Height())
		return
	}

	if !state.Transactions().Head().Compare(trx.Head()) {
		t.Errorf("the trx hashtree is invalid")
		return
	}

	if state.Amount() != amount {
		t.Errorf("the amount is invalid, expected: %d, returned: %d", amount, state.Amount())
		return
	}

	if !state.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), state.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(state)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retState, err := repository.Retrieve(*chain, height)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, state, retState)
}
