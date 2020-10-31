package expenses

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	transfer_expense "github.com/xmn-services/rod-network/domain/transfers/piastres/expenses"
)

type service struct {
	adapter     Adapter
	repository  Repository
	billService bills.Service
	lockService locks.Service
	trService   transfer_expense.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	billService bills.Service,
	lockService locks.Service,
	trService transfer_expense.Service,
) Service {
	out := service{
		adapter:     adapter,
		repository:  repository,
		billService: billService,
		lockService: lockService,
		trService:   trService,
	}

	return &out
}

// Save saves an expense
func (app *service) Save(expense Expense) error {
	hash := expense.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	content := expense.Content()
	bill := content.From()
	err = app.billService.SaveAll(bill)
	if err != nil {
		return err
	}

	if content.HasRemaining() {
		remaining := content.Remaining()
		err = app.lockService.Save(remaining)
		if err != nil {
			return err
		}
	}

	trExpense, err := app.adapter.ToTransfer(expense)
	if err != nil {
		return err
	}

	return app.trService.Save(trExpense)
}

// SaveAll saves a list of expenses
func (app *service) SaveAll(expenses []Expense) error {
	for _, oneExpense := range expenses {
		err := app.Save(oneExpense)
		if err != nil {
			return err
		}
	}

	return nil
}
