package banks

import (
	"github.com/xmn-services/rod-network/domain/memory/accounts/banks/branches"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Bank represents a bank
type Bank interface {
	entities.Immutable
	Name() string
	InstitutionNumber() string
	HasBranches() bool
	Branches() []branches.Branch
}
