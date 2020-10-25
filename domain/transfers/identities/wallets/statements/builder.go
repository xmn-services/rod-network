package statements

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	incoming         []hash.Hash
	outgoing         []hash.Hash
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		incoming:         nil,
		outgoing:         nil,
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

// WithIncoming adds incoming to the builder
func (app *builder) WithIncoming(incoming []hash.Hash) Builder {
	app.incoming = incoming
	return app
}

// WithOutgoing adds outgoing to the builder
func (app *builder) WithOutgoing(outgoing []hash.Hash) Builder {
	app.outgoing = outgoing
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Statement instance
func (app *builder) Now() (Statement, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Statement instance")
	}

	if len(app.incoming) <= 0 {
		app.incoming = nil
	}

	if len(app.outgoing) <= 0 {
		app.outgoing = nil
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	if app.incoming != nil && app.outgoing != nil {
		return createStatementWithIncomingAndOutgoing(immutable, app.incoming, app.outgoing), nil
	}

	if app.incoming != nil {
		return createStatementWithIncoming(immutable, app.incoming), nil
	}

	if app.outgoing != nil {
		return createStatementWithOutgoing(immutable, app.outgoing), nil
	}

	return createStatement(immutable), nil
}
