package cancels

import (
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	transfer_cancel "github.com/xmn-services/rod-network/domain/transfers/piastres/cancels"
)

type repository struct {
	builder           Builder
	expenseRepository expenses.Repository
	lockRepository    locks.Repository
	trRepository      transfer_cancel.Repository
}

func createRepository(
	builder Builder,
	expenseRepository expenses.Repository,
	lockRepository locks.Repository,
	trRepository transfer_cancel.Repository,
) Repository {
	out := repository{
		builder:           builder,
		expenseRepository: expenseRepository,
		lockRepository:    lockRepository,
		trRepository:      trRepository,
	}

	return &out
}

// Retrieve retrieves a cancel instance by hash
func (app *repository) Retrieve(hsh hash.Hash) (Cancel, error) {
	trCancel, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	expenseHash := trCancel.Expense()
	expense, err := app.expenseRepository.Retrieve(expenseHash)
	if err != nil {
		return nil, err
	}

	lockHash := trCancel.Lock()
	lock, err := app.lockRepository.Retrieve(lockHash)
	if err != nil {
		return nil, err
	}

	signatures := trCancel.Signatures()
	createdOn := trCancel.CreatedOn()
	return app.builder.Create().WithExpense(expense).WithLock(lock).WithSignatures(signatures).CreatedOn(createdOn).Now()
}
