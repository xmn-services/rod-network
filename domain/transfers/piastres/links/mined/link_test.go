package mined

import (
	"reflect"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestLink_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	lnk, _ := hashAdapter.FromBytes([]byte("to build the link hash..."))
	mining := "23234234523"
	createdOn := time.Now().UTC()

	link, err := NewBuilder().Create().WithHash(*hsh).WithLink(*lnk).WithMining(mining).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !link.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !link.Link().Compare(*lnk) {
		t.Errorf("the link hash is invalid")
		return
	}

	if !reflect.DeepEqual(link.Mining(), mining) {
		t.Errorf("the mining results are invalid")
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
