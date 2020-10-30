package banks

import (
	"github.com/xmn-services/rod-network/domain/memory/identities/banks/bank"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Banks represents an identity's bank
type Banks interface {
	entities.Mutable
	All() []bank.Bank
}
