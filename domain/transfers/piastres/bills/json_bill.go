package bills

import "time"

type jsonBill struct {
	Hash      string    `json:"hash"`
	Lock      string    `json:"lock"`
	Amount    uint      `json:"amount"`
	CreatedOn time.Time `json:"created_on"`
}

func createJSONBillFromBill(ins Bill) *jsonBill {
	hash := ins.Hash().String()
	lock := ins.Lock().String()
	amount := ins.Amount()
	createdOn := ins.CreatedOn()
	return createJSONBill(hash, lock, amount, createdOn)
}

func createJSONBill(
	hash string,
	lock string,
	amount uint,
	createdOn time.Time,
) *jsonBill {
	out := jsonBill{
		Hash:      hash,
		Lock:      lock,
		Amount:    amount,
		CreatedOn: createdOn,
	}

	return &out
}
