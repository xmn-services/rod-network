package blocks

import (
	"time"

	"github.com/xmn-services/rod-network/libs/hashtree"
)

type jsonBlock struct {
	Hash       string                `json:"hash"`
	Address    string                `json:"address"`
	Trx        *hashtree.JSONCompact `json:"transactions"`
	Amount     uint                  `json:"amount"`
	Additional uint                  `json:"additional"`
	CreatedOn  time.Time             `json:"created_on"`
}

func createJSONBlockFromBlock(ins Block) *jsonBlock {
	hash := ins.Hash().String()
	address := ins.Address().String()
	trx := hashtree.NewAdapter().ToJSON(ins.Transactions().Compact())
	amount := ins.Amount()
	additional := ins.Additional()
	createdOn := ins.CreatedOn()
	return createJSONBlock(hash, address, trx, amount, additional, createdOn)
}

func createJSONBlock(
	hash string,
	address string,
	trx *hashtree.JSONCompact,
	amount uint,
	additional uint,
	createdOn time.Time,
) *jsonBlock {
	out := jsonBlock{
		Hash:       hash,
		Address:    address,
		Trx:        trx,
		Amount:     amount,
		Additional: additional,
		CreatedOn:  createdOn,
	}

	return &out
}
