package application

import "github.com/xmn-services/rod-network/domain/memory/identities"

// Application represents the application
type Application interface {
	Init(name string, root string, password string, seed string) (identities.Identity, error)
	Identity(name string, password string, seed string) identities.Identity
}
