package expenses

import (
	"time"
)

type jsonExpense struct {
	Hash       string    `json:"hash"`
	From       []string  `json:"from"`
	To         string    `json:"to"`
	Signatures []string  `json:"signatures"`
	Remaining  string    `json:"remaining"`
	CreatedOn  time.Time `json:"created_on"`
}

func createJSONExpenseFromExpense(ins Expense) *jsonExpense {
	hash := ins.Hash().String()

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

	to := ins.To().String()
	createdOn := ins.CreatedOn()
	return createJSONExpense(hash, fromStrs, to, signatures, remaining, createdOn)
}

func createJSONExpense(
	hash string,
	from []string,
	to string,
	signatures []string,
	remaining string,
	createdOn time.Time,
) *jsonExpense {
	out := jsonExpense{
		Hash:       hash,
		From:       from,
		To:         to,
		Signatures: signatures,
		Remaining:  remaining,
		CreatedOn:  createdOn,
	}

	return &out
}
