package messages

import (
	"github.com/xmn-services/rod-network/domain/memory/contacts"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Messages represents messages
type Messages interface {
	All() []Message
}

// Message represents a message
type Message interface {
	entities.Immutable
	IsIncoming() bool
	Incoming() Incoming
	IsOutgoing() bool
	Outgoing() Outgoing
}

// Outgoing represents an outgoing message
type Outgoing interface {
	entities.Immutable
	To() contacts.Contact
	Subject() string
	Description() string
	HasParent() bool
	Parent() Incoming
}

// Incoming represents an incoming message
type Incoming interface {
	entities.Immutable
	From() contacts.Contact
	Subject() string
	Description() string
	HasParent() bool
	Parent() Outgoing
}

// Public represents an public message
type Public interface {
	entities.Immutable
	From() hash.Hash
	To() hash.Hash
	Subject() string
	Description() string
	HasParent() bool
	Parent() Public
}
