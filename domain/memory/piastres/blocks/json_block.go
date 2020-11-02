package blocks

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
)

// JSONBlock represents a JSON block instance
type JSONBlock struct {
	Genesis    *genesis.JSONGenesis            `json:"genesis"`
	Trx        []*transactions.JSONTransaction `json:"transactions"`
	Additional uint                            `json:"additional"`
	CreatedOn  time.Time                       `json:"created_on"`
}

func createJSONBlockFromBlock(block Block) *JSONBlock {
	genAdapter := genesis.NewAdapter()
	gen := genAdapter.ToJSON(block.Genesis())

	trxAdapter := transactions.NewAdapter()
	lst := block.Transactions()
	trx := []*transactions.JSONTransaction{}
	for _, oneTrx := range lst {
		jsTrx := trxAdapter.ToJSON(oneTrx)
		trx = append(trx, jsTrx)
	}

	additional := block.Additional()
	createdOn := block.CreatedOn()
	return createJSONBlock(gen, trx, additional, createdOn)
}

func createJSONBlock(
	gen *genesis.JSONGenesis,
	trx []*transactions.JSONTransaction,
	additional uint,
	createdOn time.Time,
) *JSONBlock {
	out := JSONBlock{
		Genesis:    gen,
		Trx:        trx,
		Additional: additional,
		CreatedOn:  createdOn,
	}

	return &out
}
