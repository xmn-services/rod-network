package public

import (
	request_public "github.com/xmn-services/rod-network/domain/memory/contacts/requests/public"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Answer represents a request public answer
type Answer interface {
	entities.Immutable
	Request() request_public.Request
	IsAccepted() bool
	Description() string
	HasPublicKey() bool
	PublicKey() public.Key
}
