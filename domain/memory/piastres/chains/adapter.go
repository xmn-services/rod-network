package chains

import (
	transfer_chains "github.com/xmn-services/rod-network/domain/transfers/piastres/chains"
)

type adapter struct {
	trBuilder transfer_chains.Builder
}

func createAdapter(
	trBuilder transfer_chains.Builder,
) Adapter {
	out := adapter{
		trBuilder: trBuilder,
	}

	return &out
}

// ToTransfer converts a chain to a transfer chain instance
func (app *adapter) ToTransfer(chain Chain) (transfer_chains.Chain, error) {
	hash := chain.Hash()
	gen := chain.Genesis().Hash()
	root := chain.Root().Hash()
	head := chain.Head().Hash()
	height := chain.Height()
	createdOn := chain.CreatedOn()
	return app.trBuilder.Create().
		WithHash(hash).
		WithGenesis(gen).
		WithRoot(root).
		WithHead(head).
		WithHeight(height).
		CreatedOn(createdOn).
		Now()
}
