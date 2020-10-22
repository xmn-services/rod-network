package chunks

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	sizeInBytes      uint
	data             *hash.Hash
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		sizeInBytes:      0,
		data:             nil,
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

// WithSizeInBytes adds a sizeInBytes to the builder
func (app *builder) WithSizeInBytes(sizeInBytes uint) Builder {
	app.sizeInBytes = sizeInBytes
	return app
}

// WithData adds data to the builder
func (app *builder) WithData(data hash.Hash) Builder {
	app.data = &data
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Chunk instance
func (app *builder) Now() (Chunk, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Chunk instance")
	}

	if app.sizeInBytes <= 0 {
		return nil, errors.New("the sizeInBytes must be greater than zero (0) in order to build a Chunk instance")
	}

	if app.data == nil {
		return nil, errors.New("the data hash is mandatory in order to build a Chunk instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createChunk(immutable, app.sizeInBytes, *app.data), nil
}
