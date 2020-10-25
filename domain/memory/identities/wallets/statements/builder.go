package statements

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/statements/entries"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	incoming         []entries.Entry
	outgoing         []entries.Entry
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		incoming:         nil,
		outgoing:         nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithIncoming adds incoming to the builder
func (app *builder) WithIncoming(incoming []entries.Entry) Builder {
	app.incoming = incoming
	return app
}

// WithOutgoing adds outgoing to the builder
func (app *builder) WithOutgoing(outgoing []entries.Entry) Builder {
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
	if len(app.incoming) <= 0 {
		app.incoming = nil
	}

	if len(app.outgoing) <= 0 {
		app.outgoing = nil
	}

	data := [][]byte{}
	if app.incoming != nil {
		for _, oneIncoming := range app.incoming {
			data = append(data, oneIncoming.Hash().Bytes())
		}
	}

	if app.outgoing != nil {
		for _, oneOutgoing := range app.outgoing {
			data = append(data, oneOutgoing.Hash().Bytes())
		}
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
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
