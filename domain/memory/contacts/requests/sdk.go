package requests

import (
	request_public "github.com/xmn-services/rod-network/domain/memory/contacts/requests/public"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Request represents a contact request
type Request interface {
	entities.Immutable
	Public() request_public.Request
	PrivateKey() encryption.PrivateKey
}
