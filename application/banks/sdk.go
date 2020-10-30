package banks

import (
	"github.com/xmn-services/rod-network/domain/memory/accounts/banks"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents a bank application
type Application interface {
	Retrieve(hash hash.Hash) (banks.Bank, error)
	RetrieveAll() ([]banks.Bank, error)
	New(name string, institutionNumber string) error
	AddBranch(bank hash.Hash, transitNumber string, address hash.Hash) error
}
