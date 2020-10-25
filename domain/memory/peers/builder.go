package peers

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/peers/peer"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter    hash.Adapter
	mutableBuilder entities.MutableBuilder
	peers          []peer.Peer
	createdOn      *time.Time
	lastUpdatedOn  *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	mutableBuilder entities.MutableBuilder,
) Builder {
	out := builder{
		hashAdapter:    hashAdapter,
		mutableBuilder: mutableBuilder,
		peers:          nil,
		createdOn:      nil,
		lastUpdatedOn:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.mutableBuilder)
}

// WithPeers add peers to the builder
func (app *builder) WithPeers(peers []peer.Peer) Builder {
	app.peers = peers
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// LastUpdatedOn adds a lastUpdatedOn time to the builder
func (app *builder) LastUpdatedOn(lastUpdatedOn time.Time) Builder {
	app.lastUpdatedOn = &lastUpdatedOn
	return app
}

// Now builds a new Peers instance
func (app *builder) Now() (Peers, error) {
	if app.peers == nil {
		app.peers = []peer.Peer{}
	}

	data := [][]byte{
		[]byte("init"),
	}

	for _, onePeer := range app.peers {
		data = append(data, onePeer.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	mutable, err := app.mutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).LastUpdatedOn(app.lastUpdatedOn).Now()
	if err != nil {
		return nil, err
	}

	return createPeers(mutable, app.peers), nil
}
