package bills

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
)

// JSONBill represents a json bill
type JSONBill struct {
	Lock      *locks.JSONLock `json:"lock"`
	Amount    uint            `json:"amount"`
	CreatedOn time.Time       `json:"created_on"`
}

func createJSONBillFromBill(ins Bill) *JSONBill {
	lockAdapter := locks.NewAdapter()
	lock := lockAdapter.ToJSON(ins.Lock())
	amount := ins.Amount()
	createdOn := ins.CreatedOn()
	return createJSONBill(lock, amount, createdOn)
}

func createJSONBill(
	lock *locks.JSONLock,
	amount uint,
	createdOn time.Time,
) *JSONBill {
	out := JSONBill{
		Lock:      lock,
		Amount:    amount,
		CreatedOn: createdOn,
	}

	return &out
}
