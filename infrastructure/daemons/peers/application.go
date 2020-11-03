package peers

import (
	"time"

	"github.com/xmn-services/rod-network/application/peers"
	domain_peers "github.com/xmn-services/rod-network/domain/memory/peers"
	"github.com/xmn-services/rod-network/domain/memory/peers/peer"
	client_peer "github.com/xmn-services/rod-network/infrastructure/clients/peers"
)

type application struct {
	remoteApplicationBuilder client_peer.Builder
	localApplication         peers.Application
	peersBuilder             domain_peers.Builder
	peersService             domain_peers.Service
	waitPeriod               time.Duration
	isStarted                bool
}

func createApplication(
	remoteApplicationBuilder client_peer.Builder,
	localApplication peers.Application,
	peersBuilder domain_peers.Builder,
	peersService domain_peers.Service,
	waitPeriod time.Duration,
) Application {
	out := application{
		remoteApplicationBuilder: remoteApplicationBuilder,
		localApplication:         localApplication,
		peersBuilder:             peersBuilder,
		peersService:             peersService,
		waitPeriod:               waitPeriod,
		isStarted:                false,
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

		peers, err := app.localApplication.Retrieve()
		if err != nil {
			return err
		}

		allPeers := []peer.Peer{}
		localPeers := peers.All()
		for _, oneLocalPeer := range localPeers {
			remoteApplication, err := app.remoteApplicationBuilder.Create().WithPeer(oneLocalPeer).Now()
			if err != nil {
				return err
			}

			remotePeers, err := remoteApplication.Retrieve()
			if err != nil {
				return err
			}

			allPeers = append(allPeers, oneLocalPeer)
			allPeers = append(allPeers, remotePeers.All()...)
		}

		updatedPeers, err := app.peersBuilder.Create().WithPeers(allPeers).Now()
		if err != nil {
			return err
		}

		err = app.peersService.Save(updatedPeers)
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
