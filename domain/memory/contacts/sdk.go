package contacts

import (
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Contact represents a contact
type Contact interface {
	entities.Immutable
	Name() string
	PublicKey() public.Key
	PrivateKey() encryption.PrivateKey
}
