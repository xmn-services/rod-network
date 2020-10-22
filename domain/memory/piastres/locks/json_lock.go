package locks

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
)

// JSONLock represents a JSON lock instance
type JSONLock struct {
	ShareHolders []*shareholders.JSONShareHolder `json:"shareholders"`
	Treeshold    uint                            `json:"treeshold"`
	Amount       uint                            `json:"amount"`
	CreatedOn    time.Time                       `json:"created_on"`
}

func createJSONLockFromLock(ins Lock) *JSONLock {
	shareHolderAdapter := shareholders.NewAdapter()
	holders := []*shareholders.JSONShareHolder{}
	lst := ins.ShareHolders()
	for _, oneShareHolder := range lst {
		holder := shareHolderAdapter.ToJSON(oneShareHolder)
		holders = append(holders, holder)
	}

	treeshold := ins.Treeshold()
	createdOn := ins.CreatedOn()
	return createJSONLock(holders, treeshold, createdOn)
}

func createJSONLock(
	shareholders []*shareholders.JSONShareHolder,
	treeshold uint,
	createdOn time.Time,
) *JSONLock {
	out := JSONLock{
		ShareHolders: shareholders,
		Treeshold:    treeshold,
		CreatedOn:    createdOn,
	}

	return &out
}
