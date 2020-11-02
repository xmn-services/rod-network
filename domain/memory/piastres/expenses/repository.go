package expenses

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	transfer_expense "github.com/xmn-services/rod-network/domain/transfers/piastres/expenses"
	"github.com/xmn-services/rod-network/libs/hash"
)

type repository struct {
	builder        Builder
	contentBuilder ContentBuilder
	billRepository bills.Repository
	trRepository   transfer_expense.Repository
}

func createRepository(
	builder Builder,
	contentBuilder ContentBuilder,
	billRepository bills.Repository,
	trRepository transfer_expense.Repository,
) Repository {
	out := repository{
		builder:        builder,
		contentBuilder: contentBuilder,
		billRepository: billRepository,
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

	billHashes := trExpense.From()
	bills, err := app.billRepository.RetrieveAll(billHashes)
	if err != nil {
		return nil, err
	}

	toHash := trExpense.To()
	to, err := app.billRepository.Retrieve(toHash)
	if err != nil {
		return nil, err
	}

	amount := to.Amount()
	lock := to.Lock()
	createdOn := trExpense.CreatedOn()
	builder := app.contentBuilder.Create().WithAmount(amount).CreatedOn(createdOn).From(bills).WithLock(lock)
	if trExpense.HasRemaining() {
		remBillHash := trExpense.Remaining()
		remaining, err := app.billRepository.Retrieve(*remBillHash)
		if err != nil {
			return nil, err
		}

		builder.WithRemaining(remaining.Lock())
	}

	content, err := builder.Now()
	if err != nil {
		return nil, err
	}

	signatures := trExpense.Signatures()
	return app.builder.Create().WithContent(content).WithSignatures(signatures).Now()
}

// RetrieveAll retrieves a list of expenses
func (app *repository) RetrieveAll(hashes []hash.Hash) ([]Expense, error) {
	out := []Expense{}
	for _, oneHash := range hashes {
		expense, err := app.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		out = append(out, expense)
	}

	return out, nil
}
