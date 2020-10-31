package transactions

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	transfer_transaction "github.com/xmn-services/rod-network/domain/transfers/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/file"
)

// CreateTransactionWithFeesForTests creates a transaction with fees for tests
func CreateTransactionWithFeesForTests(amountPubKeyInRing uint, executesOnTrigger bool, fees []expenses.Expense) Transaction {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().
		Create().
		WithFees(fees).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests(amountPubKeyInRing uint) (Repository, Service) {
	expenseRepository, expenseService := expenses.CreateRepositoryServiceForTests()
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_transaction.NewRepository(fileRepositoryService)
	trService := transfer_transaction.NewService(fileRepositoryService)
	repository := NewRepository(expenseRepository, trRepository)
	service := NewService(repository, expenseService, trService)
	return repository, service
}

// TestCompare compare two instances
func TestCompare(t *testing.T, first Transaction, second Transaction) {
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
