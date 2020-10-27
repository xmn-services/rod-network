package peers

import (
	"github.com/xmn-services/rod-network/domain/memory/peers"
	"github.com/xmn-services/rod-network/domain/memory/peers/peer"
)

type application struct {
	peerBuilder     peer.Builder
	peersBuilder    peers.Builder
	peersRepository peers.Repository
	peersService    peers.Service
}

func createApplication(
	peerBuilder peer.Builder,
	peersBuilder peers.Builder,
	peersRepository peers.Repository,
	peersService peers.Service,
) Application {
	out := application{
		peerBuilder:     peerBuilder,
		peersBuilder:    peersBuilder,
		peersRepository: peersRepository,
		peersService:    peersService,
	}

	return &out
}

// Retrieve retrieve peers
func (app *application) Retrieve() (peers.Peers, error) {
	return app.peersRepository.Retrieve()
}

// SaveClear saves a clear host
func (app *application) SaveClear(host string, port uint) error {
	builder := app.peerBuilder.Create().IsClear()
	return app.save(builder, host, port)
}

// SaveOnion saves an onion host
func (app *application) SaveOnion(host string, port uint) error {
	builder := app.peerBuilder.Create().IsOnion()
	return app.save(builder, host, port)
}

func (app *application) save(builder peer.Builder, host string, port uint) error {
	newPeer, err := builder.WithHost(host).WithPort(port).Now()
	if err != nil {
		return err
	}

	allPeers, err := app.Retrieve()
	if err != nil {
		return err
	}

	peers := allPeers.All()
	peers = append(peers, newPeer)
	updatedPeers, err := app.peersBuilder.Create().WithPeers(peers).Now()
	if err != nil {
		return err
	}

	return app.peersService.Save(updatedPeers)
}
