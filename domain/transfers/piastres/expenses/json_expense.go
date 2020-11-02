package expenses

import (
	"time"
)

type jsonExpense struct {
	Hash       string    `json:"hash"`
	Amount     uint64    `json:"amount"`
	From       []string  `json:"from"`
	Lock       string    `json:"lock"`
	Signatures []string  `json:"signatures"`
	Remaining  string    `json:"remaining"`
	CreatedOn  time.Time `json:"created_on"`
}

func createJSONExpenseFromExpense(ins Expense) *jsonExpense {
	hash := ins.Hash().String()
	amount := ins.Amount()

	fromStrs := []string{}
	from := ins.From()
	for _, oneHash := range from {
		fromStrs = append(fromStrs, oneHash.String())
	}

	signatures := []string{}
	sigs := ins.Signatures()
	for _, oneSig := range sigs {
		signatures = append(signatures, oneSig.String())
	}

	remaining := ""
	if ins.HasRemaining() {
		remaining = ins.Remaining().String()
	}

	lock := ins.Lock().String()
	createdOn := ins.CreatedOn()
	return createJSONExpense(hash, amount, fromStrs, lock, signatures, remaining, createdOn)
}

func createJSONExpense(
	hash string,
	amount uint64,
	from []string,
	lock string,
	signatures []string,
	remaining string,
	createdOn time.Time,
) *jsonExpense {
	out := jsonExpense{
		Hash:       hash,
		Amount:     amount,
		From:       from,
		Lock:       lock,
		Signatures: signatures,
		Remaining:  remaining,
		CreatedOn:  createdOn,
	}

	return &out
}
