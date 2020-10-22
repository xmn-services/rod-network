package genesis

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	transfer_genesis "github.com/xmn-services/rod-network/domain/transfers/piastres/genesis"
)

// CreateGenesisForTests creates a genesis instance for tests
func CreateGenesisForTests(blockDiffBase uint, blockDiffIncreasePerTrx float64, linkDiff uint, bill bills.Bill) Genesis {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().
		WithBlockDifficultyBase(blockDiffBase).
		WithBlockDifficultyIncreasePerTrx(blockDiffIncreasePerTrx).
		WithLinkDifficulty(linkDiff).
		WithBill(bill).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (Repository, Service) {
	fileNameWithExt := "genesis.json"
	billRepository, billService := bills.CreateRepositoryServiceForTests()
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_genesis.NewRepository(fileRepositoryService, fileNameWithExt)
	trService := transfer_genesis.NewService(fileRepositoryService, fileNameWithExt)
	repository := NewRepository(billRepository, trRepository)
	service := NewService(repository, billService, trService)
	return repository, service
}

// TestCompare compare two expense instances
func TestCompare(t *testing.T, first Genesis, second Genesis) {
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
