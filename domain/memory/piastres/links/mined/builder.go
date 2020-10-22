package mined

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/links"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	link             links.Link
	mining           string
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		link:             nil,
		mining:           "",
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithLink adds a link to the builder
func (app *builder) WithLink(link links.Link) Builder {
	app.link = link
	return app
}

// WithMining adds a mining to the builder
func (app *builder) WithMining(mining string) Builder {
	app.mining = mining
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Link instance
func (app *builder) Now() (Link, error) {
	if app.link == nil {
		return nil, errors.New("the link is mandatory in order to build a mined Link instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.link.Hash().Bytes(),
		[]byte(app.mining),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createLink(immutable, app.link, app.mining), nil
}
