package answers

import (
	answer_public "github.com/xmn-services/rod-network/domain/memory/contacts/requests/answers/public"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Answer represents a contact request answer
type Answer interface {
	entities.Immutable
	Public() answer_public.Answer
	PrivateKey() encryption.PrivateKey
}
