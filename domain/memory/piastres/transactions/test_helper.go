package transactions

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/cancels"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	transfer_transaction "github.com/xmn-services/rod-network/domain/transfers/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/file"
)

// CreateTransactionWithFeesForTests creates a transaction with fees for tests
func CreateTransactionWithFeesForTests(amountPubKeyInRing uint, executesOnTrigger bool, fees []expenses.Expense) (Transaction, signature.PrivateKey) {
	triggersOn := time.Now().UTC()
	content, err := NewContentBuilder().
		Create().
		TriggersOn(triggersOn).
		WithFees(fees).
		Now()

	if err != nil {
		panic(err)
	}

	pk := signature.NewPrivateKeyFactory().Create()
	createdOn := time.Now().UTC()
	ins, err := NewBuilder(amountPubKeyInRing).Create().WithContent(content).WithPrivateKey(pk).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins, pk
}

// CreateTransactionWithCancelForTests creates a transaction with cancel for tests
func CreateTransactionWithCancelForTests(amountPubKeyInRing uint, executesOnTrigger bool, cancel cancels.Cancel) (Transaction, signature.PrivateKey) {
	element, err := NewElementBuilder().Create().WithCancel(cancel).Now()
	if err != nil {
		panic(err)
	}

	triggersOn := time.Now().UTC()
	content, err := NewContentBuilder().
		Create().
		TriggersOn(triggersOn).
		WithElement(element).
		Now()

	if err != nil {
		panic(err)
	}

	pk := signature.NewPrivateKeyFactory().Create()
	createdOn := time.Now().UTC()
	ins, err := NewBuilder(amountPubKeyInRing).Create().WithContent(content).WithPrivateKey(pk).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins, pk
}

// CreateTransactionWithCancelAndFeesForTests creates a transaction with cancel and fees for tests
func CreateTransactionWithCancelAndFeesForTests(amountPubKeyInRing uint, executesOnTrigger bool, cancel cancels.Cancel, fees []expenses.Expense) (Transaction, signature.PrivateKey) {
	element, err := NewElementBuilder().Create().WithCancel(cancel).Now()
	if err != nil {
		panic(err)
	}

	triggersOn := time.Now().UTC()
	content, err := NewContentBuilder().
		Create().
		TriggersOn(triggersOn).
		WithElement(element).
		WithFees(fees).
		Now()

	if err != nil {
		panic(err)
	}

	pk := signature.NewPrivateKeyFactory().Create()
	createdOn := time.Now().UTC()
	ins, err := NewBuilder(amountPubKeyInRing).Create().WithContent(content).WithPrivateKey(pk).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins, pk
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests(amountPubKeyInRing uint) (Repository, Service) {
	builder := NewBuilder(amountPubKeyInRing)
	expenseRepository, expenseService := expenses.CreateRepositoryServiceForTests()
	cancelRepository, cancelService := cancels.CreateRepositoryServiceForTests()
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_transaction.NewRepository(fileRepositoryService)
	trService := transfer_transaction.NewService(fileRepositoryService)
	repository := NewRepository(builder, expenseRepository, cancelRepository, trRepository)
	service := NewService(repository, expenseService, cancelService, trService)
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
