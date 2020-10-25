package contacts

import (
	request_public "github.com/xmn-services/rod-network/domain/memory/contacts/requests/public"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents a contacts application
type Application interface {
	Request(to public.Key, subject string, description string) error
	Accept(request request_public.Request, description string) error
	Deny(request request_public.Request, description string) error
	Update(update Update) error
	Delete(hash hash.Hash) error
}

// UpdateBuilder represents an update builder
type UpdateBuilder interface {
	Create() UpdateBuilder
	WithName(name string) UpdateBuilder
	WithDescription(description string) UpdateBuilder
	Now() (Update, error)
}

// Update represents a contact update
type Update interface {
	HasName() bool
	Name() string
	HasDescription() bool
	Description() string
}
