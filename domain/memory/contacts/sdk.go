package contacts

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/contacts/requests/answers"
	transfer_contact "github.com/xmn-services/rod-network/domain/transfers/contacts"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Adapter returns the contact adapter
type Adapter interface {
	ToTransfer(contact Contact) (transfer_contact.Contact, error)
}

// Builder represents a contact builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithPublicKey(pubKey public.Key) Builder
	WithPrivateKey(pk encryption.PrivateKey) Builder
	WithAnswer(answer answers.Public) Builder
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
	Answer() answers.Public
}

// Repository represents a contact repository
type Repository interface {
	Retrieve(hash hash.Hash) (Contact, error)
}

// Service represents a contact service
type Service interface {
	Save(contact Contact) error
}
