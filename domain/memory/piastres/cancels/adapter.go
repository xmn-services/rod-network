package cancels

import (
	transfer_cancel "github.com/xmn-services/rod-network/domain/transfers/piastres/cancels"
)

type adapter struct {
	trBuilder transfer_cancel.Builder
}

func createAdapter(
	trBuilder transfer_cancel.Builder,
) Adapter {
	out := adapter{
		trBuilder: trBuilder,
	}

	return &out
}

// ToTransfer converts a Cancel instance to a transfer Cancel instance
func (app *adapter) ToTransfer(cancel Cancel) (transfer_cancel.Cancel, error) {
	hsh := cancel.Hash()
	expense := cancel.Expense().Hash()
	lock := cancel.Lock().Hash()
	signatures := cancel.Signatures()
	createdOn := cancel.CreatedOn()
	return app.trBuilder.Create().
		WithHash(hsh).
		WithExpense(expense).
		WithLock(lock).
		WithSignatures(signatures).
		CreatedOn(createdOn).
		Now()
}

// ToJSON converts a Cancel instance to a JSON instance
func (app *adapter) ToJSON(ins Cancel) *JSONCancel {
	return createJSONCancelFromCancel(ins)
}

// ToCancel converts a JSON Cancel instance to a Cancel instance
func (app *adapter) ToCancel(ins *JSONCancel) (Cancel, error) {
	return createCancelFromJSON(ins)
}
