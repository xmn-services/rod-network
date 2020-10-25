package answers

import (
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Adapter represents an answer adapter
type Adapter interface {
	ToAnswer(js []byte) (Answer, error)
	ToJSON(answer Answer) ([]byte, error)
}

// Builder represents an answer builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithPublic(hash hash.Hash) Builder
	WithPrivateKey(pk encryption.PrivateKey) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Answer, error)
}

// Answer represents a contact request answer
type Answer interface {
	entities.Immutable
	Public() hash.Hash
	PrivateKey() encryption.PrivateKey
}

// Repository represents an answer repository
type Repository interface {
	Retrieve(hash hash.Hash) (Answer, error)
}

// Service represents an answer service
type Service interface {
	Save(answer Answer) error
	Delete(answer Answer) error
}
