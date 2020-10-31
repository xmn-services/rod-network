package expenses

import (
	transfer_expense "github.com/xmn-services/rod-network/domain/transfers/piastres/expenses"
	"github.com/xmn-services/rod-network/libs/hash"
)

type adapter struct {
	trBuilder transfer_expense.Builder
}

func createAdapter(
	trBuilder transfer_expense.Builder,
) Adapter {
	out := adapter{
		trBuilder: trBuilder,
	}

	return &out
}

// ToTransfer converts an expense to a transfer expense
func (app *adapter) ToTransfer(expense Expense) (transfer_expense.Expense, error) {
	hsh := expense.Hash()
	sigs := expense.Signatures()
	content := expense.Content()
	amount := content.Amount()
	lockHash := content.Lock().Hash()
	createdOn := content.CreatedOn()

	fromHashes := []hash.Hash{}
	from := content.From()
	for _, oneFrom := range from {
		fromHashes = append(fromHashes, oneFrom.Hash())
	}

	builder := app.trBuilder.Create().
		WithHash(hsh).
		WithAmount(amount).
		From(fromHashes).
		WithLock(lockHash).
		WithSignatures(sigs).
		CreatedOn(createdOn)

	if content.HasRemaining() {
		remaining := content.Remaining().Hash()
		builder.WithRemaining(remaining)
	}

	return builder.Now()
}

// ToJSON converts an expense instance to JSON
func (app *adapter) ToJSON(ins Expense) *JSONExpense {
	return createJSONExpenseFromExpense(ins)
}

// ToExpense converts a JSON Expense instance to expense
func (app *adapter) ToExpense(ins *JSONExpense) (Expense, error) {
	return createExpenseFromJSON(ins)
}
