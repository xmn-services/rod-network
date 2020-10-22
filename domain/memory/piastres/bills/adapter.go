package bills

import (
	transfer_bill "github.com/xmn-services/rod-network/domain/transfers/piastres/bills"
)

type adapter struct {
	trBuilder transfer_bill.Builder
}

func createAdapter(
	trBuilder transfer_bill.Builder,
) Adapter {
	out := adapter{
		trBuilder: trBuilder,
	}

	return &out
}

// ToTransfer convert a bill to a transfer bill
func (app *adapter) ToTransfer(bill Bill) (transfer_bill.Bill, error) {
	hsh := bill.Hash()
	lock := bill.Lock().Hash()
	amount := bill.Amount()
	createdOn := bill.CreatedOn()
	return app.trBuilder.Create().WithHash(hsh).WithLock(lock).WithAmount(amount).CreatedOn(createdOn).Now()
}

// ToJSON converts a bill to a JSONBill instance
func (app *adapter) ToJSON(ins Bill) *JSONBill {
	return createJSONBillFromBill(ins)
}

// ToBill converts a JSONBIll to a Bill instance
func (app *adapter) ToBill(ins *JSONBill) (Bill, error) {
	return createBillFromJSON(ins)
}
