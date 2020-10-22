package application

import "github.com/xmn-services/rod-network/domain/memory/identities"

// Client represents the client application
type Client interface {
	Init(name string, root string, password string) (identities.Identity, error)
	Retrieve(name string, password string) (identities.Identity, error)
	Create(name string, password string) (identities.Identity, error)
	Save(original identities.Identity, updated identities.Identity, password string) error
	Delete(identity identities.Identity, password string) error
}

// Server represents the server application
type Server interface {
	Start() error
	Stop() error
}
