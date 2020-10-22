package cancels

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
)

// JSONCancel represents a json cancel
type JSONCancel struct {
	Expense    *expenses.JSONExpense `json:"expense"`
	Lock       *locks.JSONLock       `json:"lock"`
	Signatures []string              `json:"signatures"`
	CreatedOn  time.Time             `json:"created_on"`
}

func createJSONCancelFromCancel(ins Cancel) *JSONCancel {
	expenseAdapter := expenses.NewAdapter()
	expense := expenseAdapter.ToJSON(ins.Expense())

	lockAdapter := locks.NewAdapter()
	lock := lockAdapter.ToJSON(ins.Lock())
	createdOn := ins.CreatedOn()

	signatures := []string{}
	sigs := ins.Signatures()
	for _, oneSig := range sigs {
		signatures = append(signatures, oneSig.String())
	}

	return createJSONCancel(expense, lock, signatures, createdOn)
}

func createJSONCancel(
	expense *expenses.JSONExpense,
	lock *locks.JSONLock,
	signatures []string,
	createdOn time.Time,
) *JSONCancel {
	out := JSONCancel{
		Expense:    expense,
		Lock:       lock,
		Signatures: signatures,
		CreatedOn:  createdOn,
	}

	return &out
}
