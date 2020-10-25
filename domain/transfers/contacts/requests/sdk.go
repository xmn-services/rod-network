package requests

import (
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Adapter represents a request adapter
type Adapter interface {
	ToRequest(js []byte) (Request, error)
	ToJSON(request Request) ([]byte, error)
}

// Builder represents a request builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithPublic(public hash.Hash) Builder
	WithPrivateKey(pk encryption.PrivateKey) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Request, error)
}

// Request represents a contact request
type Request interface {
	entities.Immutable
	Public() hash.Hash
	PrivateKey() encryption.PrivateKey
}

// Repository represents a request repository
type Repository interface {
	Retrieve(hash hash.Hash) (Request, error)
}

// Service represents a request service
type Service interface {
	Save(request Request) error
	Delete(request Request) error
}
