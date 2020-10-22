package mined

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
	transfer_block_mined "github.com/xmn-services/rod-network/domain/transfers/piastres/blocks/mined"
)

type service struct {
	adapter      Adapter
	repository   Repository
	blockService blocks.Service
	trService    transfer_block_mined.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	blockService blocks.Service,
	trService transfer_block_mined.Service,
) Service {
	out := service{
		adapter:      adapter,
		repository:   repository,
		blockService: blockService,
		trService:    trService,
	}

	return &out
}

// Save saves a block
func (app *service) Save(block Block) error {
	hash := block.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	blk := block.Block()
	err = app.blockService.Save(blk)
	if err != nil {
		return err
	}

	trBlock, err := app.adapter.ToTransfer(block)
	if err != nil {
		return err
	}

	return app.trService.Save(trBlock)
}
