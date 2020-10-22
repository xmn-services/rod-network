package blocks

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	transfer_block "github.com/xmn-services/rod-network/domain/transfers/piastres/blocks"
)

// CreateBlockForTests creates a block for tests
func CreateBlockForTests(gen genesis.Genesis, additional uint, trx []transactions.Transaction) Block {
	createdOn := time.Now().UTC()
	address, _ := hash.NewAdapter().FromBytes([]byte("lets say this is the address"))
	ins, err := NewBuilder().Create().
		WithAddress(*address).
		WithGenesis(gen).
		WithAdditional(additional).
		WithTransactions(trx).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (genesis.Service, Repository, Service) {
	amountPubKeyInRing := uint(20)
	genesisRepository, genesisService := genesis.CreateRepositoryServiceForTests()
	trxRepository, trxService := transactions.CreateRepositoryServiceForTests(amountPubKeyInRing)
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_block.NewRepository(fileRepositoryService)
	trService := transfer_block.NewService(fileRepositoryService)
	repository := NewRepository(genesisRepository, trxRepository, trRepository)
	service := NewService(repository, trxService, trService)
	return genesisService, repository, service
}

// TestCompare compare two block instances
func TestCompare(t *testing.T, first Block, second Block) {
	js, err := json.Marshal(first)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = json.Unmarshal(js, second)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	reJS, err := json.Marshal(second)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if bytes.Compare(js, reJS) != 0 {
		t.Errorf("the transformed javascript is different.\n%s\n%s", js, reJS)
		return
	}

	if !first.Hash().Compare(second.Hash()) {
		t.Errorf("the instance conversion failed")
		return
	}
}
