package handshakes

import (
	"github.com/xmn-services/rod-network/domain/memory/contacts/requests/answers"
	answers_public "github.com/xmn-services/rod-network/domain/memory/contacts/requests/answers/public"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Handshake represents a contact handshake
type Handshake interface {
	entities.Immutable
	IsIncoming() bool
	Incoming() answers_public.Answer
	IsOutgoing() bool
	Outgoing() answers.Answer
}
