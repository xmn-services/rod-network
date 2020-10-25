package lists

import (
	"github.com/xmn-services/rod-network/domain/memory/identities/lists/list/contacts"
	"github.com/xmn-services/rod-network/libs/entities"
)

// List represents a contact list
type List interface {
	entities.Immutable
	Name() string
	Description() string
	HasContacts() bool
	Contacts() []contacts.Contact
}
