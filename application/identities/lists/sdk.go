package lists

import "github.com/xmn-services/rod-network/application/identities/lists/contacts"

// Application represents the contact list application
type Application interface {
	Insert(name string, description string) error
	Update(update Update) error
	Delete(name string) error
	Contacts() contacts.Application
}

// UpdateBuilder represents an update builder
type UpdateBuilder interface {
	Create() UpdateBuilder
	WithName(name string) UpdateBuilder
	WithDescription(description string) UpdateBuilder
	Now() (Update, error)
}

// Update represents an update list
type Update interface {
	HasName() bool
	Name() string
	HasDescription() bool
	Description() string
}
