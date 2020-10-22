package expenses

import (
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	transfer_expense "github.com/xmn-services/rod-network/domain/transfers/piastres/expenses"
)

type repository struct {
	builder        Builder
	contentBuilder ContentBuilder
	billRepository bills.Repository
	lockRepository locks.Repository
	trRepository   transfer_expense.Repository
}

func createRepository(
	builder Builder,
	contentBuilder ContentBuilder,
	billRepository bills.Repository,
	lockRepository locks.Repository,
	trRepository transfer_expense.Repository,
) Repository {
	out := repository{
		builder:        builder,
		contentBuilder: contentBuilder,
		billRepository: billRepository,
		lockRepository: lockRepository,
		trRepository:   trRepository,
	}

	return &out
}

// Retrieve retrieves an expense by hash
func (app *repository) Retrieve(hash hash.Hash) (Expense, error) {
	trExpense, err := app.trRepository.Retrieve(hash)
	if err != nil {
		return nil, err
	}

	billHash := trExpense.From()
	bill, err := app.billRepository.Retrieve(billHash)
	if err != nil {
		return nil, err
	}

	lockHash := trExpense.Cancel()
	lock, err := app.lockRepository.Retrieve(lockHash)
	if err != nil {
		return nil, err
	}

	amount := trExpense.Amount()
	createdOn := trExpense.CreatedOn()
	builder := app.contentBuilder.Create().WithAmount(amount).CreatedOn(createdOn).From(bill).WithCancel(lock)
	if trExpense.HasRemaining() {
		lockHash := trExpense.Remaining()
		lock, err := app.lockRepository.Retrieve(*lockHash)
		if err != nil {
			return nil, err
		}

		builder.WithRemaining(lock)
	}

	content, err := builder.Now()
	if err != nil {
		return nil, err
	}

	signatures := trExpense.Signatures()
	return app.builder.Create().WithContent(content).WithSignatures(signatures).Now()
}
