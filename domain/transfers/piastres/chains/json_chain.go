package chains

import (
	"time"
)

type jsonChain struct {
	Hash      string    `json:"hash"`
	Genesis   string    `json:"genesis"`
	Root      string    `json:"root"`
	Head      string    `json:"head"`
	Total     uint      `json:"total"`
	CreatedOn time.Time `json:"created_on"`
}

func createJSONChainFromChain(ins Chain) *jsonChain {
	hash := ins.Hash().String()
	genesis := ins.Genesis().String()
	root := ins.Root().String()
	head := ins.Head().String()
	total := ins.Total()
	createdOn := ins.CreatedOn()
	return createJSONChain(hash, genesis, root, head, total, createdOn)
}

func createJSONChain(
	hash string,
	genesis string,
	root string,
	head string,
	total uint,
	createdOn time.Time,
) *jsonChain {
	out := jsonChain{
		Hash:      hash,
		Genesis:   genesis,
		Root:      root,
		Head:      head,
		Total:     total,
		CreatedOn: createdOn,
	}

	return &out
}
