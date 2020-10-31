package statements

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/wallet/statements/entries"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Factory represents a statement factory
type Factory interface {
	Create() (Statement, error)
}

// Builder represents the statement builder
type Builder interface {
	Create() Builder
	WithIncoming(incoming []entries.Entry) Builder
	WithOutgoing(outgoing []entries.Entry) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Statement, error)
}

// Statement represents a statement
type Statement interface {
	entities.Immutable
	HasIncoming() bool
	Incoming() []entries.Entry
	HasOutgoing() bool
	Outgoing() []entries.Entry
}
