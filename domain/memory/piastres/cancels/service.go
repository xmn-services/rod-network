package cancels

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	transfer_cancel "github.com/xmn-services/rod-network/domain/transfers/piastres/cancels"
)

type service struct {
	adapter        Adapter
	repository     Repository
	expenseService expenses.Service
	lockService    locks.Service
	trService      transfer_cancel.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	expenseService expenses.Service,
	lockService locks.Service,
	trService transfer_cancel.Service,
) Service {
	out := service{
		adapter:        adapter,
		repository:     repository,
		expenseService: expenseService,
		lockService:    lockService,
		trService:      trService,
	}

	return &out
}

// Save saves a cancel instance
func (app *service) Save(cancel Cancel) error {
	hash := cancel.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	expense := cancel.Expense()
	err = app.expenseService.Save(expense)
	if err != nil {
		return err
	}

	lock := cancel.Lock()
	err = app.lockService.Save(lock)
	if err != nil {
		return err
	}

	trCancel, err := app.adapter.ToTransfer(cancel)
	if err != nil {
		return err
	}

	return app.trService.Save(trCancel)
}
