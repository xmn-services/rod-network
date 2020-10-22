package answers

import (
	"github.com/xmn-services/rod-network/domain/memory/contacts/requests"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Answer represents a contact request answer
type Answer interface {
	entities.Immutable
	Public() Public
	PrivateKey() encryption.PrivateKey
}

// Public represents a request public answer
type Public interface {
	entities.Immutable
	Request() requests.Public
	IsAccepted() bool
	PublicKey() public.Key
	Description() string
}
