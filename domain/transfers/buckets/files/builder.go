package files

import (
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	relativePath     string
	chunks           hashtree.HashTree
	amount           uint
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hash:         nil,
		relativePath: "",
		chunks:       nil,
		amount:       0,
		createdOn:    nil,
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

// WithRelativePath adds a relativePath to the builder
func (app *builder) WithRelativePath(relativePath string) Builder {
	app.relativePath = relativePath
	return app
}

// WithChunks add chunks hashtree to the builder
func (app *builder) WithChunks(chunks hashtree.HashTree) Builder {
	app.chunks = chunks
	return app
}

// WithAmount adds an amount to the builder
func (app *builder) WithAmount(amount uint) Builder {
	app.amount = amount
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new file instance
func (app *builder) Now() (File, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a File instance")
	}

	if app.relativePath == "" {
		return nil, errors.New("the relative path is mandatory in order to build a File instance")
	}

	if app.chunks == nil {
		return nil, errors.New("the chunks are mandatory in order to build a File instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount is mandatory in order to build a File instance")
	}

	length := app.chunks.Length()
	if length < int(app.amount) {
		str := fmt.Sprintf("the chunk's length (%d) cannot be smaller than the amount (%d)", length, app.amount)
		return nil, errors.New(str)
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createFile(immutable, app.relativePath, app.chunks, app.amount), nil
}
