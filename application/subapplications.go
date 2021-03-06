package application

import (
	"github.com/xmn-services/rod-network/application/genesis"
	"github.com/xmn-services/rod-network/application/miners"
	application_peers "github.com/xmn-services/rod-network/application/peers"
)

type subApplications struct {
	peerApp    application_peers.Application
	genesisApp genesis.Application
	minerApp   miners.Application
}

func createSubApplicationa(
	peerApp application_peers.Application,
	genesisApp genesis.Application,
	minerApp miners.Application,
) SubApplications {
	out := subApplications{
		peerApp:    peerApp,
		genesisApp: genesisApp,
		minerApp:   minerApp,
	}

	return &out
}

// Peers returns the peers application
func (app *subApplications) Peers() application_peers.Application {
	return app.peerApp
}

// Genesis returns the genesis application
func (app *subApplications) Genesis() genesis.Application {
	return app.genesisApp
}

// Miner returns the miner application
func (app *subApplications) Miner() miners.Application {
	return app.minerApp
}
