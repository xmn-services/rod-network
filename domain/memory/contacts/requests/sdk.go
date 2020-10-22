package requests

import (
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Request represents a contact request
type Request interface {
	entities.Immutable
	Public() Public
	PrivateKey() encryption.PrivateKey
}

// Public represents a public request
type Public interface {
	entities.Immutable
	PublicKey() public.Key
	Subject() string
	Description() string
}
