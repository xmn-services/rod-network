package transactions

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/cancels"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	transfer_transaction "github.com/xmn-services/rod-network/domain/transfers/piastres/transactions"
)

type service struct {
	adapter        Adapter
	repository     Repository
	expenseService expenses.Service
	cancelService  cancels.Service
	trService      transfer_transaction.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	expenseService expenses.Service,
	cancelService cancels.Service,
	trService transfer_transaction.Service,
) Service {
	out := service{
		adapter:        adapter,
		repository:     repository,
		expenseService: expenseService,
		cancelService:  cancelService,
		trService:      trService,
	}

	return &out
}

// Save saves a transaction
func (app *service) Save(trx Transaction) error {
	hash := trx.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	content := trx.Content()
	if content.HasFees() {
		fees := content.Fees()
		err = app.expenseService.Save(fees)
		if err != nil {
			return err
		}
	}

	if content.IsExpense() {
		expense := content.Expense()
		err = app.expenseService.Save(expense)
		if err != nil {
			return err
		}
	}

	if content.IsCancel() {
		cancel := content.Cancel()
		err = app.cancelService.Save(cancel)
		if err != nil {
			return err
		}
	}

	trTrx, err := app.adapter.ToTransfer(trx)
	if err != nil {
		return err
	}

	return app.trService.Save(trTrx)
}

// SaveAll saves all transactions
func (app *service) SaveAll(trx []Transaction) error {
	for _, oneTrx := range trx {
		err := app.Save(oneTrx)
		if err != nil {
			return err
		}
	}

	return nil
}
