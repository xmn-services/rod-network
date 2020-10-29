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
	DeleteBranch(bank hash.Hash, transitNumber string, address hash.Hash) error
	Update(hash hash.Hash, update Update) error
	Merge(hashes []hash.Hash) error
	Delete(hash hash.Hash) error
}

// Update represents an update
type Update interface {
	HasName() bool
	Name() string
	HasInstitutionNumber() bool
	InstitutionNumber() string
}
