package blocks

import (
	"github.com/xmn-services/rod-network/libs/hashtree"
	transfer_block "github.com/xmn-services/rod-network/domain/transfers/piastres/blocks"
)

type adapter struct {
	hashTreeBuilder hashtree.Builder
	trBuilder       transfer_block.Builder
}

func createAdapter(
	hashTreeBuilder hashtree.Builder,
	trBuilder transfer_block.Builder,
) Adapter {
	out := adapter{
		hashTreeBuilder: hashTreeBuilder,
		trBuilder:       trBuilder,
	}

	return &out
}

// ToTransfer converts the block to a transfer block
func (app *adapter) ToTransfer(block Block) (transfer_block.Block, error) {
	trx := block.Transactions()
	blocks := [][]byte{}
	for _, oneTrx := range trx {
		blocks = append(blocks, oneTrx.Hash().Bytes())
	}

	trxHashtree, err := app.hashTreeBuilder.Create().WithBlocks(blocks).Now()
	if err != nil {
		return nil, err
	}

	hsh := block.Hash()
	address := block.Address()
	additional := block.Additional()
	amount := uint(len(trx))
	createdOn := block.CreatedOn()
	return app.trBuilder.Create().
		WithHash(hsh).
		WithAddress(address).
		WithAdditional(additional).
		WithTransactions(trxHashtree).
		WithAmount(amount).
		CreatedOn(createdOn).
		Now()
}

// ToJSON converts a block to a JSON block
func (app *adapter) ToJSON(block Block) *JSONBlock {
	return createJSONBlockFromBlock(block)
}

// ToBlock converts a JSON block to a block
func (app *adapter) ToBlock(ins *JSONBlock) (Block, error) {
	return createBlockFromJSON(ins)
}
