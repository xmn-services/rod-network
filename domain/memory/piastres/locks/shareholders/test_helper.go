package shareholders

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
	transfer_lock_shareholder "github.com/xmn-services/rod-network/domain/transfers/piastres/locks/shareholders"
)

// CreateShareHolderForTests creates a shareholder for tests
func CreateShareHolderForTests(power uint, key hash.Hash) ShareHolder {
	createdOn := time.Now().UTC()
	holder, err := NewBuilder().Create().WithPower(power).WithKey(key).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return holder
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (Repository, Service) {
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_lock_shareholder.NewRepository(fileRepositoryService)
	trService := transfer_lock_shareholder.NewService(fileRepositoryService)
	repository := NewRepository(trRepository)
	service := NewService(repository, trService)
	return repository, service
}

// TestCompare compare two expense instances
func TestCompare(t *testing.T, first ShareHolder, second ShareHolder) {
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
