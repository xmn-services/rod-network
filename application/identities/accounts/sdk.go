package accounts

import (
	"github.com/xmn-services/rod-network/domain/memory/accounts"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents an address application
type Application interface {
	Retrieve(hash hash.Hash) (accounts.Account, error)
	RetrieveAll() ([]accounts.Account, error)
	New(number string, bank hash.Hash) error
	Delete(hash hash.Hash) error
}
