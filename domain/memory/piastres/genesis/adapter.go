package genesis

import (
	transfer_genesis "github.com/xmn-services/rod-network/domain/transfers/piastres/genesis"
)

type adapter struct {
	trBuilder transfer_genesis.Builder
}

func createAdapter(
	trBuilder transfer_genesis.Builder,
) Adapter {
	out := adapter{
		trBuilder: trBuilder,
	}

	return &out
}

// ToTransfer converts a genesis to a transfer genesis instance
func (app *adapter) ToTransfer(genesis Genesis) (transfer_genesis.Genesis, error) {
	hsh := genesis.Hash()
	bill := genesis.Bill().Hash()
	diff := genesis.Difficulty()
	blockDiff := diff.Block()
	blockDiffBase := blockDiff.Base()
	blockDiffIncr := blockDiff.IncreasePerTrx()
	linkDiff := diff.Link()
	createdOn := genesis.CreatedOn()

	return app.trBuilder.Create().
		WithHash(hsh).
		WithBlockDifficultyBase(blockDiffBase).
		WithBlockDifficultyIncreasePerTrx(blockDiffIncr).
		WithLinkDifficulty(linkDiff).
		WithBill(bill).
		CreatedOn(createdOn).
		Now()

}

// ToJSON converts a genesis to a JSON instance
func (app *adapter) ToJSON(genesis Genesis) *JSONGenesis {
	return createJSONGenesisFromGenesis(genesis)
}

// ToGenesis converts a JSON Genesis to a Genesis instance
func (app *adapter) ToGenesis(ins *JSONGenesis) (Genesis, error) {
	return createGenesisFromJSON(ins)
}
