package chains

import (
	"time"

	application_chain "github.com/xmn-services/rod-network/application/chains"
	application_peer "github.com/xmn-services/rod-network/application/peers"
	"github.com/xmn-services/rod-network/domain/memory/piastres/chains"
	client_chain "github.com/xmn-services/rod-network/infrastructure/clients/chains"
	client_link "github.com/xmn-services/rod-network/infrastructure/clients/chains/links"
)

type application struct {
	chainApp              application_chain.Application
	peerApp               application_peer.Application
	remoteChainAppBuilder client_chain.Builder
	remoteLinkAppBuilder  client_link.Builder
	waitPeriod            time.Duration
	isStarted             bool
}

func createApplication(
	chainApp application_chain.Application,
	peerApp application_peer.Application,
	remoteChainAppBuilder client_chain.Builder,
	remoteLinkAppBuilder client_link.Builder,
	waitPeriod time.Duration,
) Application {
	out := application{
		chainApp:              chainApp,
		peerApp:               peerApp,
		remoteChainAppBuilder: remoteChainAppBuilder,
		remoteLinkAppBuilder:  remoteLinkAppBuilder,
		waitPeriod:            waitPeriod,
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

		// retrieve the local chain:
		localChain, err := app.chainApp.Retrieve()
		if err != nil {
			return err
		}

		// retrieve the peers:
		localPeers, err := app.peerApp.Retrieve()
		if err != nil {
			return err
		}

		biggestDiff := 0
		var biggestChain chains.Chain
		var biggestChainApp application_chain.Application
		allPeers := localPeers.All()
		for _, onePeer := range allPeers {
			remoteChainApp, err := app.remoteChainAppBuilder.Create().WithPeer(onePeer).Now()
			if err != nil {
				return err
			}

			remoteChain, err := remoteChainApp.Retrieve()
			if err != nil {
				return err
			}

			diffTrx := int(remoteChain.Total() - localChain.Total())
			if biggestDiff < diffTrx {
				biggestDiff = diffTrx
				biggestChain = remoteChain
				biggestChainApp = remoteChainApp
			}
		}

		// if there is no chain in the network more advanced, continue:
		if biggestChain == nil {
			continue
		}

		// update the chain:
		localIndex := int(localChain.Height())
		diffHeight := int(biggestChain.Height()) - localIndex
		for i := 0; i < diffHeight; i++ {
			chainIndex := localIndex + i
			remoteChainAtIndex, err := biggestChainApp.RetrieveAtIndex(uint(chainIndex))
			if err != nil {
				return err
			}

			remoteHead := remoteChainAtIndex.Head()
			err = app.chainApp.Upgrade(remoteHead)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

// Stop stops the application
func (app *application) Stop() error {
	app.isStarted = true
	return nil
}
