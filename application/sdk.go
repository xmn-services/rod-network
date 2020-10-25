package application

import (
	application_identities "github.com/xmn-services/rod-network/application/identities"
	application_peers "github.com/xmn-services/rod-network/application/peers"
)

// Application represents the application
type Application interface {
	Init(name string, root string, password string, seed string) error
	Identity(name string, password string, seed string) application_identities.Application
	Peer() application_peers.Application
}
