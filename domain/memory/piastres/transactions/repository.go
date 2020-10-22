package transactions

import (
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/cancels"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	transfer_transaction "github.com/xmn-services/rod-network/domain/transfers/piastres/transactions"
)

type repository struct {
	builder           Builder
	contentBuilder    ContentBuilder
	expenseRepository expenses.Repository
	cancelRepository  cancels.Repository
	trRepository      transfer_transaction.Repository
}

func createRepository(
	builder Builder,
	contentBuilder ContentBuilder,
	expenseRepository expenses.Repository,
	cancelRepository cancels.Repository,
	trRepository transfer_transaction.Repository,
) Repository {
	out := repository{
		builder:           builder,
		contentBuilder:    contentBuilder,
		expenseRepository: expenseRepository,
		cancelRepository:  cancelRepository,
		trRepository:      trRepository,
	}

	return &out
}

// Retrieve retrieves a transaction by hash
func (app *repository) Retrieve(hsh hash.Hash) (Transaction, error) {
	trTrx, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	triggersOn := trTrx.TriggersOn()
	executesOnTrigger := trTrx.ExecutesOnTrigger()
	builder := app.contentBuilder.Create().TriggersOn(triggersOn)
	if trTrx.HasFees() {
		feesHash := trTrx.Fees()
		fees, err := app.expenseRepository.Retrieve(*feesHash)
		if err != nil {
			return nil, err
		}

		builder.WithFees(fees)
	}

	if trTrx.IsExpense() {
		expenseHash := trTrx.Expense()
		expense, err := app.expenseRepository.Retrieve(*expenseHash)
		if err != nil {
			return nil, err
		}

		builder.WithExpense(expense)
	}

	if trTrx.IsCancel() {
		cancelHash := trTrx.Cancel()
		cancel, err := app.cancelRepository.Retrieve(*cancelHash)
		if err != nil {
			return nil, err
		}

		builder.WithCancel(cancel)
	}

	if executesOnTrigger {
		builder.ExecutesOnTrigger()
	}

	content, err := builder.Now()
	if err != nil {
		return nil, err
	}

	signature := trTrx.Signature()
	createdOn := trTrx.CreatedOn()
	return app.builder.Create().WithContent(content).WithSignature(signature).CreatedOn(createdOn).Now()
}

// RetrieveAll retrieves all trx from hashes
func (app *repository) RetrieveAll(hashes []hash.Hash) ([]Transaction, error) {
	out := []Transaction{}
	for _, oneHash := range hashes {
		trx, err := app.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		out = append(out, trx)
	}

	return out, nil
}
