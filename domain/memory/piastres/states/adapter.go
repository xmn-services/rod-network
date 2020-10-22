package states

import (
	"github.com/xmn-services/rod-network/libs/hashtree"
	transfer_state "github.com/xmn-services/rod-network/domain/transfers/piastres/states"
)

type adapter struct {
	hashTreeBuilder hashtree.Builder
	trBuilder       transfer_state.Builder
}

func createAdapter(
	hashTreeBuilder hashtree.Builder,
	trBuilder transfer_state.Builder,
) Adapter {
	out := adapter{
		hashTreeBuilder: hashTreeBuilder,
		trBuilder:       trBuilder,
	}

	return &out
}

// ToTransfer converts a state to a transfer state instance
func (app *adapter) ToTransfer(state State) (transfer_state.State, error) {
	hsh := state.Hash()
	prev := state.Previous()
	height := state.Height()
	createdOn := state.CreatedOn()

	blocks := [][]byte{}
	trx := state.Transactions()
	for _, oneTrx := range trx {
		blocks = append(blocks, oneTrx.Hash().Bytes())
	}

	ht, err := app.hashTreeBuilder.Create().WithBlocks(blocks).Now()
	if err != nil {
		return nil, err
	}

	return app.trBuilder.Create().
		WithHash(hsh).
		WithPrevious(prev).
		WithHeight(height).
		WithTransactions(ht).
		WithHeight(height).
		CreatedOn(createdOn).
		Now()
}
