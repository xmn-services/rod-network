package buckets

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/identities"
	"github.com/xmn-services/rod-network/domain/memory/peers"
	client_bucket "github.com/xmn-services/rod-network/infrastructure/clients/identities/buckets"
)

type application struct {
	peersRepository        peers.Repository
	identityRepository     identities.Repository
	identityService        identities.Service
	remoteBucketAppBuilder client_bucket.Builder
	waitPeriod             time.Duration
	name                   string
	password               string
	seed                   string
	isStarted              bool
}

func createApplication(
	peersRepository peers.Repository,
	identityRepository identities.Repository,
	identityService identities.Service,
	remoteBucketAppBuilder client_bucket.Builder,
	waitPeriod time.Duration,
	name string,
	password string,
	seed string,
) Application {
	out := application{
		peersRepository:        peersRepository,
		identityRepository:     identityRepository,
		identityService:        identityService,
		remoteBucketAppBuilder: remoteBucketAppBuilder,
		waitPeriod:             waitPeriod,
		name:                   name,
		password:               password,
		seed:                   seed,
		isStarted:              false,
	}

	return &out
}

// Start starts the application
func (app *application) Start() error {
	app.isStarted = true

	for {
		// wait period:
		time.Sleep(app.waitPeriod)

		// if the application is not started, continue:
		if !app.isStarted {
			continue
		}

		// retrieve the identity:
		identity, err := app.identityRepository.Retrieve(app.name, app.password, app.seed)
		if err != nil {
			return err
		}

		// retrieve the peers:
		peers, err := app.peersRepository.Retrieve()
		if err != nil {
			return err
		}

		// build the client bucket:
		removeBucketApp, err := app.remoteBucketAppBuilder.Create().
			WithName(app.name).
			WithPassword(app.password).
			WithSeed(app.seed).
			WithPeers(peers).
			Now()

		if err != nil {
			return err
		}

		followBuckets := identity.Buckets().Follows()
		for _, oneBucketHash := range followBuckets {
			bucket, err := removeBucketApp.Retrieve(oneBucketHash)
			if err != nil {
				return err
			}

			err = identity.Buckets().Follow(bucket)
			if err != nil {
				return err
			}
		}

		// save the identity:
		err = app.identityService.Update(
			identity.Hash(),
			identity,
			app.password,
			app.password,
		)

		if err != nil {
			return err
		}
	}
}

// Stop stops the application
func (app *application) Stop() error {
	app.isStarted = true
	return nil
}
