package application

import (
	"github.com/xmn-services/rod-network/application/genesis"
	application_identities "github.com/xmn-services/rod-network/application/identities"
	application_peers "github.com/xmn-services/rod-network/application/peers"
)

// Application represents the application
type Application interface {
	Current() Current
	Sub() SubApplications
}

// Current represents the current application
type Current interface {
	NewIdentity(name string, password string, seed string, root string) error
	Authenticate(name string, seed string, password string) (application_identities.Application, error)
}

// SubApplications represents the sub applications
type SubApplications interface {
	Peers() application_peers.Application
	Genesis() genesis.Application
}
