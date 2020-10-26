package peer

import (
	"errors"
)

type builder struct {
	host    string
	port    uint
	isClear bool
	isOnion bool
}

func createBuilder() Builder {
	out := builder{
		host:    "",
		port:    0,
		isClear: false,
		isOnion: false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
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

// Now builds a new Peer instance
func (app *builder) Now() (Peer, error) {
	if app.host == "" {
		return nil, errors.New("the host is mandatory in order to build a Peer instance")
	}

	if app.port == 0 {
		return nil, errors.New("the port is mandatory in order to build a Peer instance")
	}

	if app.isClear {
		return createPeerWithClear(app.host, app.port), nil
	}

	if app.isOnion {
		return createPeerWithOnion(app.host, app.port), nil
	}

	return nil, errors.New("the Peer is invalid")
}
