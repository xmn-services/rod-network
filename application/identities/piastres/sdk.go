package piastres

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
)

// Application represents a piastre application
type Application interface {
	Current() Current
}

// Current represents the current application
type Current interface {
	Bucket(absolutePath string, fees []Fee) error
}

// Fee represents a fee
type Fee interface {
	Amount() uint64
	Lock() locks.Lock
}
