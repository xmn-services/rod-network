package application

import (
	application_identities "github.com/xmn-services/rod-network/application/identities"
	application_peers "github.com/xmn-services/rod-network/application/peers"
	"github.com/xmn-services/rod-network/domain/memory/identities"
)

type application struct {
	identityAppBuilder application_identities.Builder
	peerApp            application_peers.Application
	identityBuilder    identities.Builder
	identityRepository identities.Repository
	identityService    identities.Service
}

func createApplication(
	identityAppBuilder application_identities.Builder,
	peerApp application_peers.Application,
	identityBuilder identities.Builder,
	identityRepository identities.Repository,
	identityService identities.Service,
) Application {
	out := application{
		identityAppBuilder: identityAppBuilder,
		peerApp:            peerApp,
		identityBuilder:    identityBuilder,
		identityRepository: identityRepository,
		identityService:    identityService,
	}

	return &out
}

// Peers returns the peers application
func (app *application) Peers() application_peers.Application {
	return app.peerApp
}

// Init initializes the network
func (app *application) Init(name string, root string, password string, seed string) error {
	return nil
}

// NewIdentity saves a new identity
func (app *application) NewIdentity(name string, password string, seed string, root string) error {
	identity, err := app.identityBuilder.Create().WithName(name).WithSeed(seed).WithRoot(root).Now()
	if err != nil {
		return err
	}

	return app.identityService.Insert(identity, password)
}

// Authenticate authenticates an identity
func (app *application) Authenticate(name string, seed string, password string) (application_identities.Application, error) {
	identity, err := app.identityRepository.Retrieve(name, seed, password)
	if err != nil {
		return nil, err
	}

	return app.identityAppBuilder.Create().WithIdentity(identity).Now()
}
