package transfers

import (
	"github.com/xmn-services/rod-network/domain/memory/addresses"
	"github.com/xmn-services/rod-network/domain/memory/exchanges/reserves/orders/transfers/deposits"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Transfer represents a transfer
type Transfer interface {
	entities.Immutable
	IsPerson() bool
	Person() addresses.Address
	IsDeposit() bool
	Deposit() deposits.Deposit
}
