package peers

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/peers/peer"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	mutableBuilder entities.MutableBuilder
	hash           *hash.Hash
	peers          []peer.Peer
	createdOn      *time.Time
	lastUpdatedOn  *time.Time
}

func createBuilder(
	mutableBuilder entities.MutableBuilder,
) Builder {
	out := builder{
		mutableBuilder: mutableBuilder,
		hash:           nil,
		peers:          nil,
		createdOn:      nil,
		lastUpdatedOn:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.mutableBuilder)
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.hash = &hash
	return app
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
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Peers instance")
	}

	if app.peers == nil {
		app.peers = []peer.Peer{}
	}

	mutable, err := app.mutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).LastUpdatedOn(app.lastUpdatedOn).Now()
	if err != nil {
		return nil, err
	}

	return createPeers(mutable, app.peers), nil
}
