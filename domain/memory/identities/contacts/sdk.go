package contacts

import (
	"github.com/xmn-services/rod-network/domain/memory/contacts"
	"github.com/xmn-services/rod-network/domain/memory/contacts/requests"
	"github.com/xmn-services/rod-network/domain/memory/contacts/requests/answers"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Contact represents a contact
type Contact interface {
	entities.Immutable
	IsAccepted() bool
	Accepted() contacts.Contact
	IsPending() bool
	Pending() requests.Request
	IsDenied() bool
	Denied() answers.Public
}
