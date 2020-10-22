package mined

import "time"

type jsonBlock struct {
	Hash      string    `json:"hash"`
	Block     string    `json:"block"`
	Mining    string    `json:"mining"`
	CreatedOn time.Time `json:"created_on"`
}

func createJSONBlockFromBlock(ins Block) *jsonBlock {
	hash := ins.Hash().String()
	block := ins.Block().String()
	mining := ins.Mining()
	createdOn := ins.CreatedOn()
	return createJSONBlock(hash, block, mining, createdOn)
}

func createJSONBlock(
	hash string,
	block string,
	mining string,
	createdOn time.Time,
) *jsonBlock {
	out := jsonBlock{
		Hash:      hash,
		Block:     block,
		Mining:    mining,
		CreatedOn: createdOn,
	}

	return &out
}
