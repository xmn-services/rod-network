package owners

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills/owners/balances"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills/owners/owner"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills/owners/statements"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Owners represents owners
type Owners interface {
	entities.Immutable
	All() []owner.Owner
	Balances() balances.Balances
	Statement() statements.Statement
}
