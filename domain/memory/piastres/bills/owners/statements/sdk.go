package statements

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills/owners/statements/entries"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Statement represents a statement
type Statement interface {
	entities.Immutable
	HasIncoming() bool
	Incoming() []entries.Entry
	HasOutgoing() bool
	Outgoing() []entries.Entry
}
