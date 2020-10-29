package orders

import (
	"github.com/xmn-services/rod-network/domain/memory/exchanges/reserves"
	"github.com/xmn-services/rod-network/domain/memory/exchanges/reserves/orders/transfers"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Order represents a reserve order
type Order interface {
	entities.Immutable
	Content() Content
	Signature() signature.RingSignature
}

// Content represents an order content
type Content interface {
	entities.Immutable
	Reserve() reserves.Reserve
	Transfer() transfers.Transfer
	Lock() locks.Lock
}
