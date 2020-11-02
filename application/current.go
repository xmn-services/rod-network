package application

import (
	application_identities "github.com/xmn-services/rod-network/application/identities"
	"github.com/xmn-services/rod-network/domain/memory/identities"
)

type current struct {
	identityAppBuilder application_identities.Builder
	identityBuilder    identities.Builder
	identityRepository identities.Repository
	identityService    identities.Service
}

func createCurrent(
	identityAppBuilder application_identities.Builder,
	identityBuilder identities.Builder,
	identityRepository identities.Repository,
	identityService identities.Service,
) Current {
	out := current{
		identityAppBuilder: identityAppBuilder,
		identityBuilder:    identityBuilder,
		identityRepository: identityRepository,
		identityService:    identityService,
	}

	return &out
}

// NewIdentity saves a new identity
func (app *current) NewIdentity(name string, password string, seed string, root string) error {
	identity, err := app.identityBuilder.Create().WithName(name).WithSeed(seed).WithRoot(root).Now()
	if err != nil {
		return err
	}

	return app.identityService.Insert(identity, password)
}

// Authenticate authenticates an identity
func (app *current) Authenticate(name string, seed string, password string) (application_identities.Application, error) {
	identity, err := app.identityRepository.Retrieve(name, seed, password)
	if err != nil {
		return nil, err
	}

	iName := identity.Name()
	iSeed := identity.Seed()
	return app.identityAppBuilder.Create().
		WithName(iName).
		WithSeed(iSeed).
		WithPassword(password).
		Now()
}
