package banks

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/accounts/banks/branches"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Builder represents a bank builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithInstitutionNumber(insNumber string) Builder
	WithBranches(branches []branches.Branch) Builder
	CreatedOn(createdOn time.Time) Builder
	LastUpdatedOn(lastUpdatedOn time.Time) Builder
	Now() (Bank, error)
}

// Bank represents a bank
type Bank interface {
	entities.Mutable
	Name() string
	InstitutionNumber() string
	HasBranches() bool
	Branches() []branches.Branch
}

// Repository represents a bank repository
type Repository interface {
	Retrieve(hash hash.Hash) (Bank, error)
	RetrieveAll() ([]Bank, error)
}

// Service represents a bank service
type Service interface {
	Insert(bank Bank) error
	Update(original Bank, updated Bank) error
}
