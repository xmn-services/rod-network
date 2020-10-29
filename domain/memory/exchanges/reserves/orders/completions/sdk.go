package completions

import (
	"github.com/xmn-services/rod-network/domain/memory/exchanges/reserves/orders"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Completion represents a completion
type Completion interface {
	entities.Immutable
	Content() Content
	Signature() signature.RingSignature
}

// Content represents an order content
type Content interface {
	entities.Immutable
	Order() orders.Order
	Expense() expenses.Expense
}
