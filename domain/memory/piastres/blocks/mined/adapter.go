package mined

import (
	transfer_block_mined "github.com/xmn-services/rod-network/domain/transfers/piastres/blocks/mined"
)

type adapter struct {
	trBuilder transfer_block_mined.Builder
}

func createAdapter(
	trBuilder transfer_block_mined.Builder,
) Adapter {
	out := adapter{
		trBuilder: trBuilder,
	}

	return &out
}

// ToTransfer converts a mined block to a transfer mined block
func (app *adapter) ToTransfer(block Block) (transfer_block_mined.Block, error) {
	hsh := block.Hash()
	blk := block.Block().Hash()
	mining := block.Mining()
	createdOn := block.CreatedOn()
	return app.trBuilder.Create().WithHash(hsh).WithBlock(blk).WithMining(mining).CreatedOn(createdOn).Now()
}
