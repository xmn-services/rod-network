package locks

import (
	"time"

	"github.com/xmn-services/rod-network/libs/hashtree"
)

type jsonLock struct {
	Hash         string                `json:"hash"`
	ShareHolders *hashtree.JSONCompact `json:"shareholders"`
	Treeshold    uint                  `json:"treeshold"`
	Amount       uint                  `json:"amount"`
	CreatedOn    time.Time             `json:"created_on"`
}

func createJSONLockFromLock(ins Lock) *jsonLock {
	hash := ins.Hash().String()
	holders := hashtree.NewAdapter().ToJSON(ins.ShareHolders().Compact())
	treeshold := ins.Treeshold()
	amount := ins.Amount()
	createdOn := ins.CreatedOn()
	return createJSONLock(hash, holders, treeshold, amount, createdOn)
}

func createJSONLock(
	hash string,
	shareholders *hashtree.JSONCompact,
	treeshold uint,
	amount uint,
	createdOn time.Time,
) *jsonLock {
	out := jsonLock{
		Hash:         hash,
		ShareHolders: shareholders,
		Treeshold:    treeshold,
		Amount:       amount,
		CreatedOn:    createdOn,
	}

	return &out
}
