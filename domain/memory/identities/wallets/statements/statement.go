package statements

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/statements/entries"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type statement struct {
	immutable entities.Immutable
	incoming  []entries.Entry
	outgoing  []entries.Entry
}

func createStatement(
	immutable entities.Immutable,
) Statement {
	return createStatementInternally(immutable, nil, nil)
}

func createStatementWithIncoming(
	immutable entities.Immutable,
	incoming []entries.Entry,
) Statement {
	return createStatementInternally(immutable, incoming, nil)
}

func createStatementWithOutgoing(
	immutable entities.Immutable,
	outgoing []entries.Entry,
) Statement {
	return createStatementInternally(immutable, nil, outgoing)
}

func createStatementWithIncomingAndOutgoing(
	immutable entities.Immutable,
	incoming []entries.Entry,
	outgoing []entries.Entry,
) Statement {
	return createStatementInternally(immutable, incoming, outgoing)
}

func createStatementInternally(
	immutable entities.Immutable,
	incoming []entries.Entry,
	outgoing []entries.Entry,
) Statement {
	out := statement{
		immutable: immutable,
		incoming:  incoming,
		outgoing:  outgoing,
	}

	return &out
}

// Hash returns the hash
func (obj *statement) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// HasIncoming returns true if there is incoming, false otherwise
func (obj *statement) HasIncoming() bool {
	return obj.incoming != nil
}

// Incoming returns the incoming, if any
func (obj *statement) Incoming() []entries.Entry {
	return obj.incoming
}

// HasOutgoing returns true if there is outgoing, false otherwise
func (obj *statement) HasOutgoing() bool {
	return obj.outgoing != nil
}

// Outgoing returns the outgoing, if any
func (obj *statement) Outgoing() []entries.Entry {
	return obj.outgoing
}

// CreatedOn returns the creation time
func (obj *statement) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
