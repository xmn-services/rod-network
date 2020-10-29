package deposits

import (
	"github.com/xmn-services/rod-network/domain/memory/buckets/informations"
	"github.com/xmn-services/rod-network/domain/memory/accounts"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Deposit represents a deposit
type Deposit interface {
	entities.Immutable
	Account() accounts.Account
	Receipt() informations.Information
}
