package links

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	previousLink     *hash.Hash
	next             *hash.Hash
	index            uint
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		previousLink:     nil,
		next:             nil,
		index:            0,
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

// WithPreviousLink adds a previousLink hash to the builder
func (app *builder) WithPreviousLink(prevLink hash.Hash) Builder {
	app.previousLink = &prevLink
	return app
}

// WithNext adds a next hash to the builder
func (app *builder) WithNext(next hash.Hash) Builder {
	app.next = &next
	return app
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.index = index
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
		return nil, errors.New("the hash is mandatory in order to build a Link instance")
	}

	if app.previousLink == nil {
		return nil, errors.New("the previousLink hash is mandatory in order to build a Link instance")
	}

	if app.next == nil {
		return nil, errors.New("the next hash is mandatory in order to build a Link instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createLink(immutable, *app.previousLink, *app.next, app.index), nil
}
