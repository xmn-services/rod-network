package owners

import "github.com/xmn-services/rod-network/domain/memory/piastres/bills/owners/owner"

// Owners represents owners
type Owners interface {
	All() []owner.Owner
}
