package piastres

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
)

// Application represents a piastre application
type Application interface {
	Spend(spend Spend) error
	SpendWithFees(spend Spend, fees Spend) error
	Cancel(expense expenses.Expense, lock locks.Lock, signatures []signature.RingSignature) error
}

// Spend represents a spend
type Spend interface {
	Amount() uint
	Lock() locks.Lock
	TriggersOn() time.Time
	Name() string
	Description() string
}
