package mined

import (
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
	transfer_block_mined "github.com/xmn-services/rod-network/domain/transfers/piastres/blocks/mined"
)

type repository struct {
	builder         Builder
	blockRepository blocks.Repository
	trRepository    transfer_block_mined.Repository
}

func createRepository(
	builder Builder,
	blockRepository blocks.Repository,
	trRepository transfer_block_mined.Repository,
) Repository {
	out := repository{
		builder:         builder,
		blockRepository: blockRepository,
		trRepository:    trRepository,
	}

	return &out
}

// Retrieve retrieves a block by hash
func (app *repository) Retrieve(hsh hash.Hash) (Block, error) {
	trBlock, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	blockHash := trBlock.Block()
	block, err := app.blockRepository.Retrieve(blockHash)
	if err != nil {
		return nil, err
	}

	mining := trBlock.Mining()
	createdOn := trBlock.CreatedOn()
	return app.builder.Create().WithBlock(block).WithMining(mining).CreatedOn(createdOn).Now()
}
