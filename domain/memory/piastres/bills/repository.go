package bills

import (
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	transfer_bill "github.com/xmn-services/rod-network/domain/transfers/piastres/bills"
)

type repository struct {
	builder        Builder
	lockRepository locks.Repository
	trRepository   transfer_bill.Repository
}

func createRepository(
	builder Builder,
	lockRepository locks.Repository,
	trRepository transfer_bill.Repository,
) Repository {
	out := repository{
		builder:        builder,
		lockRepository: lockRepository,
		trRepository:   trRepository,
	}

	return &out
}

// Retrieve retrieves a bill by hash
func (app *repository) Retrieve(hsh hash.Hash) (Bill, error) {
	trBill, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	lockHash := trBill.Lock()
	lock, err := app.lockRepository.Retrieve(lockHash)
	if err != nil {
		return nil, err
	}

	amount := trBill.Amount()
	createdOn := trBill.CreatedOn()
	return app.builder.Create().WithLock(lock).WithAmount(amount).CreatedOn(createdOn).Now()
}
