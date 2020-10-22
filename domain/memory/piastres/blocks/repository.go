package blocks

import (
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	transfer_block "github.com/xmn-services/rod-network/domain/transfers/piastres/blocks"
)

type repository struct {
	builder           Builder
	genesisRepository genesis.Repository
	trxRepository     transactions.Repository
	trRepository      transfer_block.Repository
}

func createRepository(
	builder Builder,
	genesisRepository genesis.Repository,
	trxRepository transactions.Repository,
	trRepository transfer_block.Repository,
) Repository {
	out := repository{
		builder:           builder,
		genesisRepository: genesisRepository,
		trxRepository:     trxRepository,
		trRepository:      trRepository,
	}

	return &out
}

// Retrieve retrieves a block by hash
func (app *repository) Retrieve(hsh hash.Hash) (Block, error) {
	trBlock, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	gen, err := app.genesisRepository.Retrieve()
	if err != nil {
		return nil, err
	}

	trxHashes := []hash.Hash{}
	amountTrx := trBlock.Amount()
	leaves := trBlock.Transactions().Parent().BlockLeaves().Leaves()
	for i := 0; i < int(amountTrx); i++ {
		trxHashes = append(trxHashes, leaves[i].Head())
	}

	trx, err := app.trxRepository.RetrieveAll(trxHashes)
	if err != nil {
		return nil, err
	}

	address := trBlock.Address()
	additional := trBlock.Additional()
	createdOn := trBlock.CreatedOn()
	return app.builder.Create().
		WithAddress(address).
		WithGenesis(gen).
		WithAdditional(additional).
		WithTransactions(trx).
		CreatedOn(createdOn).
		Now()
}
