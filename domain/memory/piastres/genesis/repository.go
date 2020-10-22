package genesis

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	transfer_genesis "github.com/xmn-services/rod-network/domain/transfers/piastres/genesis"
)

type repository struct {
	billRepository bills.Repository
	trRepository   transfer_genesis.Repository
	builder        Builder
}

func createRepository(
	builder Builder,
	billRepository bills.Repository,
	trRepository transfer_genesis.Repository,
) Repository {
	out := repository{
		builder:        builder,
		billRepository: billRepository,
		trRepository:   trRepository,
	}

	return &out
}

// Retrieve retrieves a genesis instance
func (app *repository) Retrieve() (Genesis, error) {
	trGen, err := app.trRepository.Retrieve()
	if err != nil {
		return nil, err
	}

	billHash := trGen.Bill()
	bill, err := app.billRepository.Retrieve(billHash)
	if err != nil {
		return nil, err
	}

	blockDiffBase := trGen.BlockDifficultyBase()
	blockDiffIncreasePerTrx := trGen.BlockDifficultyIncreasePerTrx()
	linkDiff := trGen.LinkDifficulty()
	createdOn := trGen.CreatedOn()
	return app.builder.Create().
		WithBill(bill).
		WithBlockDifficultyBase(blockDiffBase).
		WithBlockDifficultyIncreasePerTrx(blockDiffIncreasePerTrx).
		WithLinkDifficulty(linkDiff).
		CreatedOn(createdOn).
		Now()
}
