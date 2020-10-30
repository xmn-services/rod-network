package states

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	transfer_state "github.com/xmn-services/rod-network/domain/transfers/piastres/states"
	"github.com/xmn-services/rod-network/libs/hash"
)

type repository struct {
	trxRepository transactions.Repository
	trRepository  transfer_state.Repository
	builder       Builder
}

func createRepository(
	trxRepository transactions.Repository,
	trRepository transfer_state.Repository,
	builder Builder,
) Repository {
	out := repository{
		trxRepository: trxRepository,
		trRepository:  trRepository,
		builder:       builder,
	}

	return &out
}

// Retrieve retrieves a state by chain and height
func (app *repository) Retrieve(chain hash.Hash, height uint) (State, error) {
	trState, err := app.trRepository.Retrieve(chain, height)
	if err != nil {
		return nil, err
	}

	amount := trState.Amount()
	trxHashes := []hash.Hash{}
	leaves := trState.Transactions().Parent().BlockLeaves().Leaves()
	for i := 0; i < int(amount); i++ {
		trxHashes = append(trxHashes, leaves[i].Head())
	}

	trx, err := app.trxRepository.RetrieveAll(trxHashes)
	if err != nil {
		return nil, err
	}

	prev := trState.Previous()
	createdOn := trState.CreatedOn()
	return app.builder.Create().
		WithChain(chain).
		WithHeight(height).
		WithPrevious(prev).
		WithTransactions(trx).
		CreatedOn(createdOn).
		Now()
}
