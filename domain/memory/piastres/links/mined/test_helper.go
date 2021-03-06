package mined

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	"github.com/xmn-services/rod-network/domain/memory/piastres/links"
	transfer_mined_link "github.com/xmn-services/rod-network/domain/transfers/piastres/links/mined"
)

// CreateLinkForTests creates a link instance for tests
func CreateLinkForTests(link links.Link, mining string) Link {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithLink(link).WithMining(mining).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (genesis.Service, Repository, Service) {
	genService, linkRepository, linkService := links.CreateRepositoryServiceForTests()
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_mined_link.NewRepository(fileRepositoryService)
	trService := transfer_mined_link.NewService(fileRepositoryService)
	repository := NewRepository(linkRepository, trRepository)
	service := NewService(repository, linkService, trService)
	return genService, repository, service
}

// TestCompare compare two expense instances
func TestCompare(t *testing.T, first Link, second Link) {
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
