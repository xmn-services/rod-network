package chains

import (
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestChain_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	genesis, _ := hashAdapter.FromBytes([]byte("to build the genesis hash..."))
	root, _ := hashAdapter.FromBytes([]byte("to build the root hash..."))
	head, _ := hashAdapter.FromBytes([]byte("to build the head hash..."))
	total := uint(1556)
	createdOn := time.Now().UTC()

	chain, err := NewBuilder().Create().WithHash(*hsh).WithGenesis(*genesis).WithRoot(*root).WithHead(*head).WithTotal(total).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !chain.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !chain.Genesis().Compare(*genesis) {
		t.Errorf("the genesis hash is invalid")
		return
	}

	if !chain.Root().Compare(*root) {
		t.Errorf("the root hash is invalid")
		return
	}

	if !chain.Head().Compare(*head) {
		t.Errorf("the head hash is invalid")
		return
	}

	if chain.Total() != total {
		t.Errorf("the total is invalid, expected: %d, returned: %d", chain.Total(), total)
		return
	}

	if !chain.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), chain.CreatedOn().String())
		return
	}

	// repository and service:
	fileNameWithExt := "chain.json"
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService, fileNameWithExt)
	service := NewService(fileRepositoryService, fileNameWithExt)

	err = service.Save(chain)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retChain, err := repository.Retrieve()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, chain, retChain)
}
