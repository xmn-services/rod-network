package genesis

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
)

// JSONGenesis represents a JSON genesis
type JSONGenesis struct {
	Bill                          *bills.JSONBill `json:"bill"`
	LinkDifficulty                uint            `json:"link_difficulty"`
	BlockDifficultyBase           uint            `json:"block_difficulty_base"`
	BlockDifficultyIncreasePerTrx float64         `json:"block_difficulty_increase_per_transaction"`
	CreatedOn                     time.Time       `json:"created_on"`
}

func createJSONGenesisFromGenesis(gen Genesis) *JSONGenesis {
	billsAdapter := bills.NewAdapter()
	bill := billsAdapter.ToJSON(gen.Bill())
	difficulty := gen.Difficulty()
	linkDifficulty := difficulty.Link()
	blockDifficulty := difficulty.Block()
	blockDifficultyBase := blockDifficulty.Base()
	blockDifficultyIncreasePerTrx := blockDifficulty.IncreasePerTrx()
	createdOn := gen.CreatedOn()
	return createJSONGenesis(
		bill,
		linkDifficulty,
		blockDifficultyBase,
		blockDifficultyIncreasePerTrx,
		createdOn,
	)
}

func createJSONGenesis(
	bill *bills.JSONBill,
	linkDifficulty uint,
	blockDifficultyBase uint,
	blockDifficultyIncreasePerTrx float64,
	createdOn time.Time,
) *JSONGenesis {
	out := JSONGenesis{
		Bill:                          bill,
		LinkDifficulty:                linkDifficulty,
		BlockDifficultyBase:           blockDifficultyBase,
		BlockDifficultyIncreasePerTrx: blockDifficultyIncreasePerTrx,
		CreatedOn:                     createdOn,
	}

	return &out
}
