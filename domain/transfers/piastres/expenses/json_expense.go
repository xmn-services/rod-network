package expenses

import (
	"time"
)

type jsonExpense struct {
	Hash       string    `json:"hash"`
	Amount     uint      `json:"amount"`
	From       string    `json:"from"`
	Cancel     string    `json:"cancel"`
	Signatures []string  `json:"signatures"`
	Remaining  string    `json:"remaining"`
	CreatedOn  time.Time `json:"created_on"`
}

func createJSONExpenseFromExpense(ins Expense) *jsonExpense {
	hash := ins.Hash().String()
	amount := ins.Amount()
	from := ins.From().String()
	cancel := ins.Cancel().String()

	signatures := []string{}
	sigs := ins.Signatures()
	for _, oneSig := range sigs {
		signatures = append(signatures, oneSig.String())
	}

	remaining := ""
	if ins.HasRemaining() {
		remaining = ins.Remaining().String()
	}

	createdOn := ins.CreatedOn()
	return createJSONExpense(hash, amount, from, cancel, signatures, remaining, createdOn)
}

func createJSONExpense(
	hash string,
	amount uint,
	from string,
	cancel string,
	signatures []string,
	remaining string,
	createdOn time.Time,
) *jsonExpense {
	out := jsonExpense{
		Hash:       hash,
		Amount:     amount,
		From:       from,
		Cancel:     cancel,
		Signatures: signatures,
		Remaining:  remaining,
		CreatedOn:  createdOn,
	}

	return &out
}
