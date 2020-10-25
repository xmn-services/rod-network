package public

import (
	"hash"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Adapter represents a public request adapter
type Adapter interface {
	ToRequest(js []byte) (Request, error)
	ToJSON(request Request) ([]byte, error)
}

// Builder represents a request builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithPublicKey(pubKey public.Key) Builder
	WithSubject(subject string) Builder
	WithDescription(description string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Request, error)
}

// Request represents a public request
type Request interface {
	entities.Immutable
	PublicKey() public.Key
	Subject() string
	HasDescription() bool
	Description() string
}

// Repository represents a public request repository
type Repository interface {
	Retrieve(hash hash.Hash) (Request, error)
}

// Service represents a public request service
type Service interface {
	Save(request Request) error
	Delete(request Request) error
}
