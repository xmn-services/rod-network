package links

import (
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestLink_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	previousLink, _ := hashAdapter.FromBytes([]byte("to build the previous link hash..."))
	next, _ := hashAdapter.FromBytes([]byte("to build the next hash..."))
	index := uint(32)
	createdOn := time.Now().UTC()

	link, err := NewBuilder().Create().WithHash(*hsh).WithPreviousLink(*previousLink).WithNext(*next).WithIndex(index).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !link.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !link.PreviousLink().Compare(*previousLink) {
		t.Errorf("the previousLink hash is invalid")
		return
	}

	if !link.Next().Compare(*next) {
		t.Errorf("the next hash is invalid")
		return
	}

	if link.Index() != index {
		t.Errorf("the index is invalid, expected: %d, returned: %d", index, link.Index())
		return
	}

	if !link.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), link.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(link)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retLink, err := repository.Retrieve(link.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, link, retLink)
}
