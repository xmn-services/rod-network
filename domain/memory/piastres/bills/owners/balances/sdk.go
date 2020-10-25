package balances

import "github.com/xmn-services/rod-network/domain/memory/piastres/bills/owners/balances/balance"

// Balances represents balances
type Balances interface {
	All() []balance.Balance
}
