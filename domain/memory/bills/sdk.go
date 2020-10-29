package bills

import (
	"github.com/xmn-services/rod-network/domain/memory/bills/currencies"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Bill represents a fiat bill
type Bill interface {
	entities.Immutable
	Currency() currencies.Currency
	Amount() uint
}
