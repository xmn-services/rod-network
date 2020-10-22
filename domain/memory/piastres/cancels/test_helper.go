package cancels

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	transfer_cancel "github.com/xmn-services/rod-network/domain/transfers/piastres/cancels"
)

// CreateCancelForTests creates a cancel instance for tests
func CreateCancelForTests(expense expenses.Expense, lock locks.Lock, sigs []signature.RingSignature) Cancel {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithExpense(expense).WithLock(lock).WithSignatures(sigs).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (Repository, Service) {
	expenseRepository, expenseService := expenses.CreateRepositoryServiceForTests()
	lockRepository, lockService := locks.CreateRepositoryServiceForTests()
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_cancel.NewRepository(fileRepositoryService)
	trService := transfer_cancel.NewService(fileRepositoryService)
	repository := NewRepository(expenseRepository, lockRepository, trRepository)
	service := NewService(repository, expenseService, lockService, trService)
	return repository, service
}

// TestCompare compare two cancel instances
func TestCompare(t *testing.T, first Cancel, second Cancel) {
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
