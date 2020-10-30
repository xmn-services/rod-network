package banks

import (
	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents a bank application
type Application interface {
	DeleteBranch(bank hash.Hash, branch hash.Hash) error
	Update(hash hash.Hash, update Update) error
	Merge(master hash.Hash, slaves []hash.Hash) error
	Delete(hash hash.Hash) error
}

// UpdateBuilder represents an update builder
type UpdateBuilder interface {
	Create() UpdateBuilder
	WithName(name string) UpdateBuilder
	WithInstitutionNumber(insNumber string) UpdateBuilder
	Now() (Update, error)
}

// Update represents an update
type Update interface {
	HasName() bool
	Name() string
	HasInstitutionNumber() bool
	InstitutionNumber() string
}
