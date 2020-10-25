package application

import (
	application_identities "github.com/xmn-services/rod-network/application/identities"
	application_peers "github.com/xmn-services/rod-network/application/peers"
	"github.com/xmn-services/rod-network/domain/memory/identities"
)

// Application represents the application
type Application interface {
	Init(name string, root string, password string, seed string) (identities.Identity, error)
	Identity(name string, password string, seed string) application_identities.Application
	Peer() application_peers.Application
}
