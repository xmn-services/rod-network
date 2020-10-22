package locks

import (
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
	transfer_lock "github.com/xmn-services/rod-network/domain/transfers/piastres/locks"
)

type repository struct {
	shareHolderRepository shareholders.Repository
	trRepository          transfer_lock.Repository
	builder               Builder
}

func createRepository(
	shareHolderRepository shareholders.Repository,
	trRepository transfer_lock.Repository,
	builder Builder,
) Repository {
	out := repository{
		shareHolderRepository: shareHolderRepository,
		trRepository:          trRepository,
		builder:               builder,
	}

	return &out
}

// Retrieve retrieves a lock by hash
func (app *repository) Retrieve(hsh hash.Hash) (Lock, error) {
	trLock, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	amount := trLock.Amount()
	holderHashes := []hash.Hash{}
	leaves := trLock.ShareHolders().Parent().BlockLeaves().Leaves()
	for i := 0; i < int(amount); i++ {
		holderHashes = append(holderHashes, leaves[i].Head())
	}

	holders, err := app.shareHolderRepository.RetrieveAll(holderHashes)
	if err != nil {
		return nil, err
	}

	treeshold := trLock.Treeshold()
	createdOn := trLock.CreatedOn()
	return app.builder.Create().WithShareHolders(holders).WithTreeshold(treeshold).CreatedOn(createdOn).Now()
}
