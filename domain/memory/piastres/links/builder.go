package links

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	prevLink         *hash.Hash
	next             blocks.Block
	index            uint
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		prevLink:         nil,
		next:             nil,
		index:            0,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithPreviousLink adds a previous link hash to the builder
func (app *builder) WithPreviousLink(prevLink hash.Hash) Builder {
	app.prevLink = &prevLink
	return app
}

// WithNext adds a next block to the builder
func (app *builder) WithNext(next blocks.Block) Builder {
	app.next = next
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
	if app.prevLink == nil {
		return nil, errors.New("the previousLink hash is mandatory in order to build a Link instance")
	}

	if app.next == nil {
		return nil, errors.New("the next block is mandatory in order to build a Link instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.prevLink.Bytes(),
		app.next.Hash().Bytes(),
		[]byte(strconv.Itoa(int(app.index))),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createLink(immutable, *app.prevLink, app.next, app.index), nil
}
