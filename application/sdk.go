package application

import (
	application_identities "github.com/xmn-services/rod-network/application/identities"
	application_peers "github.com/xmn-services/rod-network/application/peers"
)

// Application represents the application
type Application interface {
	Peers() application_peers.Application
	Init(name string, root string, password string, seed string) error
	NewIdentity(name string, password string, seed string, root string) error
	Authenticate(name string, password string, seed string) application_identities.Application
}
