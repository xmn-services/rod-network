package blocks

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
)

// JSONBlock represents a JSON block instance
type JSONBlock struct {
	Address    string                          `json:"address"`
	Genesis    *genesis.JSONGenesis            `json:"genesis"`
	Trx        []*transactions.JSONTransaction `json:"transactions"`
	Additional uint                            `json:"additional"`
	CreatedOn  time.Time                       `json:"created_on"`
}

func createJSONBlockFromBlock(block Block) *JSONBlock {
	address := block.Address().String()

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
	return createJSONBlock(address, gen, trx, additional, createdOn)
}

func createJSONBlock(
	address string,
	gen *genesis.JSONGenesis,
	trx []*transactions.JSONTransaction,
	additional uint,
	createdOn time.Time,
) *JSONBlock {
	out := JSONBlock{
		Address:    address,
		Genesis:    gen,
		Trx:        trx,
		Additional: additional,
		CreatedOn:  createdOn,
	}

	return &out
}
