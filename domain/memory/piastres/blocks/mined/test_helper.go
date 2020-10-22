package mined

import (
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
	transfer_block_mined "github.com/xmn-services/rod-network/domain/transfers/piastres/blocks/mined"
)

// CreateBlockForTests creates a mined block instance for tests
func CreateBlockForTests(blk blocks.Block, mining string) Block {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithBlock(blk).WithMining(mining).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (Repository, Service) {
	_, blockRepository, blockService := blocks.CreateRepositoryServiceForTests()
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_block_mined.NewRepository(fileRepositoryService)
	trService := transfer_block_mined.NewService(fileRepositoryService)
	repository := NewRepository(blockRepository, trRepository)
	service := NewService(repository, blockService, trService)
	return repository, service
}
