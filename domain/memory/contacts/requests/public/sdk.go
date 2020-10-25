package public

import (
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Request represents a public request
type Request interface {
	entities.Immutable
	PublicKey() public.Key
	Subject() string
	HasDescription() bool
	Description() string
}
