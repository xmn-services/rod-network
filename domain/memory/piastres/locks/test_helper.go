package locks

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
	transfer_lock "github.com/xmn-services/rod-network/domain/transfers/piastres/locks"
)

// CreateLockForTests create lock for tests
func CreateLockForTests(holders []shareholders.ShareHolder, treeshold uint) Lock {
	createdOn := time.Now().UTC()
	lock, err := NewBuilder().Create().WithShareHolders(holders).WithTreeshold(treeshold).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return lock
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (Repository, Service) {
	holderRepository, holderService := shareholders.CreateRepositoryServiceForTests()
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_lock.NewRepository(fileRepositoryService)
	trService := transfer_lock.NewService(fileRepositoryService)
	repository := NewRepository(holderRepository, trRepository)
	service := NewService(repository, holderService, trService)
	return repository, service
}

// TestCompare compare two expense instances
func TestCompare(t *testing.T, first Lock, second Lock) {
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
