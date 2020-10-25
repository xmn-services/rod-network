package peer

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	host             string
	port             uint
	isClear          bool
	isOnion          bool
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		host:             "",
		port:             0,
		isClear:          false,
		isOnion:          false,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.immutableBuilder)
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.hash = &hash
	return app
}

// WithHost adds a host to the builder
func (app *builder) WithHost(host string) Builder {
	app.host = host
	return app
}

// WithPort adds a port to the builder
func (app *builder) WithPort(port uint) Builder {
	app.port = port
	return app
}

// IsClear flags the builder as clear
func (app *builder) IsClear() Builder {
	app.isClear = true
	return app
}

// IsOnion flags the builder as onion
func (app *builder) IsOnion() Builder {
	app.isOnion = true
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Peer instance
func (app *builder) Now() (Peer, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Peer instance")
	}

	if app.port == 0 {
		return nil, errors.New("the port is mandatory in order to build a Peer instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	if app.isClear {
		return createPeerWithClear(immutable, app.host, app.port), nil
	}

	if app.isOnion {
		return createPeerWithOnion(immutable, app.host, app.port), nil
	}

	return nil, errors.New("the Peer is invalid")
}
