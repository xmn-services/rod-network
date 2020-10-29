package accounts

import (
	"github.com/xmn-services/rod-network/domain/memory/accounts/banks"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Account represets a brank account
type Account interface {
	entities.Immutable
	Bank() banks.Bank
	Number() string
}
