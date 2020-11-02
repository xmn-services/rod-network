package expenses

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	transfer_expense "github.com/xmn-services/rod-network/domain/transfers/piastres/expenses"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/file"
)

// CreateExpenseForTests creates an expense instance for tests
func CreateExpenseForTests(content Content, signatures []signature.RingSignature) Expense {
	ins, err := NewBuilder().Create().WithContent(content).WithSignatures(signatures).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// CreateContentForTests creates an expense instance for tests
func CreateContentForTests(amount uint64, from []bills.Bill, lock locks.Lock) Content {
	createdOn := time.Now().UTC()
	ins, err := NewContentBuilder().Create().WithAmount(amount).From(from).WithLock(lock).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// CreateContentWithRemainingForTests creates an expense instance with remaining for tests
func CreateContentWithRemainingForTests(amount uint64, from []bills.Bill, remaining locks.Lock) Content {
	createdOn := time.Now().UTC()
	ins, err := NewContentBuilder().Create().WithAmount(amount).From(from).WithRemaining(remaining).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (Repository, Service) {
	billRepository, billService := bills.CreateRepositoryServiceForTests()
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_expense.NewRepository(fileRepositoryService)
	trService := transfer_expense.NewService(fileRepositoryService)
	repository := NewRepository(billRepository, trRepository)
	service := NewService(repository, billService, trService)
	return repository, service
}

// TestCompare compare two expense instances
func TestCompare(t *testing.T, first Expense, second Expense) {
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
