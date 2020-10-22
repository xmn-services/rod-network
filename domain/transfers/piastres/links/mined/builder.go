package mined

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	link             *hash.Hash
	mining           string
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		link:             nil,
		mining:           "",
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

// WithLink adds a link hash to the builder
func (app *builder) WithLink(link hash.Hash) Builder {
	app.link = &link
	return app
}

// WithMining adds mining results to the builder
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
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a mined Link instance")
	}

	if app.link == nil {
		return nil, errors.New("the link hash is mandatory in order to build a mined Link instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createLink(immutable, *app.link, app.mining), nil

}
