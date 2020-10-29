package exchanges

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/bills"
	"github.com/xmn-services/rod-network/domain/memory/exchanges/methods"
	piastre_bills "github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Exchange represents a fiat currency to piastre exchange
type Exchange interface {
	entities.Immutable
	Content() Content
	Signature() signature.RingSignature
}

// Content represents an exchange content
type Content interface {
	entities.Immutable
	Method() methods.Method
	From() piastre_bills.Bill
	To() bills.Bill
	ExpireOn() time.Time
	HasNotes() bool
	Notes() string
}
