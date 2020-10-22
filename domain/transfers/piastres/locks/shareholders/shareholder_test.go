package shareholders

import (
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestShareHolder_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	key, _ := hashAdapter.FromBytes([]byte("to build the pubKey hash..."))
	power := uint(67)
	createdOn := time.Now().UTC()

	shareHolder, err := NewBuilder().Create().WithHash(*hsh).WithKey(*key).WithPower(power).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !shareHolder.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !shareHolder.Key().Compare(*key) {
		t.Errorf("the key hash is invalid")
		return
	}

	if shareHolder.Power() != power {
		t.Errorf("the power is invalid, expected: %d, returned: %d", power, shareHolder.Power())
		return
	}

	if !shareHolder.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), shareHolder.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(shareHolder)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retShareHolder, err := repository.Retrieve(shareHolder.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, shareHolder, retShareHolder)
}
