package genesis

import (
	"time"
)

type jsonGenesis struct {
	Hash                    string    `json:"hash"`
	Bill                    string    `json:"bill"`
	BlockDiffBase           uint      `json:"block_difficulty_base"`
	BlockDiffIncreasePerTrx float64   `json:"block_difficulty_increase_per_trx"`
	LinkDiff                uint      `json:"link_difficulty"`
	CreatedOn               time.Time `json:"created_on"`
}

func createJSONGenesisFromGenesis(ins Genesis) *jsonGenesis {
	hash := ins.Hash().String()
	bill := ins.Bill().String()
	blockDiffBase := ins.BlockDifficultyBase()
	blockDiffIncreasePerTrx := ins.BlockDifficultyIncreasePerTrx()
	linkDiff := ins.LinkDifficulty()
	createdOn := ins.CreatedOn()
	return createJSONGenesis(hash, bill, blockDiffBase, blockDiffIncreasePerTrx, linkDiff, createdOn)
}

func createJSONGenesis(
	hash string,
	bill string,
	blockDiffBase uint,
	blockDiffIncreasePerTrx float64,
	linkDiff uint,
	createdOn time.Time,
) *jsonGenesis {
	out := jsonGenesis{
		Hash:                    hash,
		Bill:                    bill,
		BlockDiffBase:           blockDiffBase,
		BlockDiffIncreasePerTrx: blockDiffIncreasePerTrx,
		LinkDiff:                linkDiff,
		CreatedOn:               createdOn,
	}

	return &out
}
