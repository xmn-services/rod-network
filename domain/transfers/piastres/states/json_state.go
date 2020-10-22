package states

import (
	"time"

	"github.com/xmn-services/rod-network/libs/hashtree"
)

type jsonState struct {
	Hash      string                `json:"hash"`
	Chain     string                `json:"chain"`
	Prev      string                `json:"previous"`
	Height    uint                  `json:"height"`
	Trx       *hashtree.JSONCompact `json:"transactions"`
	Amount    uint                  `json:"amount"`
	CreatedOn time.Time             `json:"created_on"`
}

func createJSONStateFromState(ins State) *jsonState {
	hash := ins.Hash().String()
	chain := ins.Chain().String()
	prev := ins.Previous().String()
	height := ins.Height()
	trx := hashtree.NewAdapter().ToJSON(ins.Transactions().Compact())
	amount := ins.Amount()
	createdOn := ins.CreatedOn()
	return createJSONState(hash, chain, prev, height, trx, amount, createdOn)
}

func createJSONState(
	hash string,
	chain string,
	prev string,
	height uint,
	trx *hashtree.JSONCompact,
	amount uint,
	createdOn time.Time,
) *jsonState {
	out := jsonState{
		Hash:      hash,
		Chain:     chain,
		Prev:      prev,
		Height:    height,
		Trx:       trx,
		Amount:    amount,
		CreatedOn: createdOn,
	}

	return &out
}
