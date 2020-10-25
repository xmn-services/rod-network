package public

import (
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Adapter represents a public answer adapter
type Adapter interface {
	ToAnswer(js []byte) (Answer, error)
	ToJSON(answer Answer) ([]byte, error)
}

// Builder represents an answer builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithRequest(request hash.Hash) Builder
	WithDescription(description string) Builder
	WithPublicKey(pubKey public.Key) Builder
	Now() (Answer, error)
}

// Answer represents a request public answer
type Answer interface {
	entities.Immutable
	Request() hash.Hash
	HasPublicKey() bool
	PublicKey() public.Key
	HasDescription() bool
	Description() string
}

// Repository represents a public answer repository
type Repository interface {
	Retrieve(hash hash.Hash) (Answer, error)
}

// Service represents a public answer service
type Service interface {
	Save(answer Answer) error
	Delete(answer Answer) error
}
