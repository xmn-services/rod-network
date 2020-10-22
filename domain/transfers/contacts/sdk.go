package contacts

import (
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Builder represents a contact builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithPublicKey(pubKey public.Key) Builder
	WithPrivateKey(pk encryption.PrivateKey) Builder
	WithAnswer(answer hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Contact, error)
}

// Contact represents a contact
type Contact interface {
	entities.Immutable
	Name() string
	Description() string
	PublicKey() public.Key
	PrivateKey() encryption.PrivateKey
	Answer() hash.Hash
}

// Repository represents a contact repository
type Repository interface {
	Retrieve(hash hash.Hash) (Contact, error)
}

// Service represents a contact service
type Service interface {
	Save(contact Contact) error
}
