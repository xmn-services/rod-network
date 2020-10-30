package bank

import (
	"github.com/xmn-services/rod-network/domain/memory/accounts/banks"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Bank represents a bank
type Bank interface {
	entities.Mutable
	Bank() banks.Bank
	HasSlaves() bool
	Slaves() []hash.Hash
}
