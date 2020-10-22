package bills

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	transfer_bill "github.com/xmn-services/rod-network/domain/transfers/piastres/bills"
)

// CreateBillForTests creates a bill for tests
func CreateBillForTests(lock locks.Lock, amount uint) Bill {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithLock(lock).WithAmount(amount).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (Repository, Service) {
	lockRepository, lockService := locks.CreateRepositoryServiceForTests()
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_bill.NewRepository(fileRepositoryService)
	trService := transfer_bill.NewService(fileRepositoryService)
	repository := NewRepository(lockRepository, trRepository)
	service := NewService(repository, lockService, trService)
	return repository, service
}

// TestCompare compare two bill instances
func TestCompare(t *testing.T, first Bill, second Bill) {
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
