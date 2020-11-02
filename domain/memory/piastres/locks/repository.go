package locks

import (
	transfer_lock "github.com/xmn-services/rod-network/domain/transfers/piastres/locks"
	"github.com/xmn-services/rod-network/libs/hash"
)

type repository struct {
	trRepository transfer_lock.Repository
	builder      Builder
}

func createRepository(
	trRepository transfer_lock.Repository,
	builder Builder,
) Repository {
	out := repository{
		trRepository: trRepository,
		builder:      builder,
	}

	return &out
}

// Retrieve retrieves a lock by hash
func (app *repository) Retrieve(hsh hash.Hash) (Lock, error) {
	trLock, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	publicKeys := trLock.PublicKeys()
	createdOn := trLock.CreatedOn()
	return app.builder.Create().WithPublicKeys(publicKeys).CreatedOn(createdOn).Now()
}
