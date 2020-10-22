package bills

import (
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestBill_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	lock, _ := hashAdapter.FromBytes([]byte("to build the lock hash..."))
	amount := uint(56)
	createdOn := time.Now().UTC()

	bill, err := NewBuilder().Create().WithHash(*hsh).WithLock(*lock).WithAmount(amount).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bill.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !bill.Lock().Compare(*lock) {
		t.Errorf("the lock hash is invalid")
		return
	}

	if bill.Amount() != amount {
		t.Errorf("the amount was expected to be %d, %d returned", amount, bill.Amount())
		return
	}

	if !bill.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), bill.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(bill)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBill, err := repository.Retrieve(bill.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, bill, retBill)
}
