package cancels

import (
	"time"
)

type jsonCancel struct {
	Hash       string    `json:"hash"`
	Expense    string    `json:"expense"`
	Lock       string    `json:"lock"`
	Signatures []string  `json:"signatures"`
	CreatedOn  time.Time `json:"created_on"`
}

func createJSONCancelFromCancel(ins Cancel) *jsonCancel {
	hash := ins.Hash().String()
	expense := ins.Expense().String()
	lock := ins.Lock().String()
	createdOn := ins.CreatedOn()

	signatures := []string{}
	sigs := ins.Signatures()
	for _, oneSig := range sigs {
		signatures = append(signatures, oneSig.String())
	}

	return createJSONCancel(hash, expense, lock, signatures, createdOn)
}

func createJSONCancel(
	hash string,
	expense string,
	lock string,
	signatures []string,
	createdOn time.Time,
) *jsonCancel {
	out := jsonCancel{
		Hash:       hash,
		Expense:    expense,
		Lock:       lock,
		Signatures: signatures,
		CreatedOn:  createdOn,
	}

	return &out
}
