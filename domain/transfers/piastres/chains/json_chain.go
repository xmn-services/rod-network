package chains

import (
	"time"
)

type jsonChain struct {
	Hash      string    `json:"hash"`
	Genesis   string    `json:"genesis"`
	Root      string    `json:"root"`
	Head      string    `json:"head"`
	Height    uint      `json:"height"`
	CreatedOn time.Time `json:"created_on"`
}

func createJSONChainFromChain(ins Chain) *jsonChain {
	hash := ins.Hash().String()
	genesis := ins.Genesis().String()
	root := ins.Root().String()
	head := ins.Head().String()
	height := ins.Height()
	createdOn := ins.CreatedOn()
	return createJSONChain(hash, genesis, root, head, height, createdOn)
}

func createJSONChain(
	hash string,
	genesis string,
	root string,
	head string,
	height uint,
	createdOn time.Time,
) *jsonChain {
	out := jsonChain{
		Hash:      hash,
		Genesis:   genesis,
		Root:      root,
		Head:      head,
		Height:    height,
		CreatedOn: createdOn,
	}

	return &out
}
