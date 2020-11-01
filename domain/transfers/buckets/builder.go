package buckets

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
	files            hashtree.HashTree
	amount           uint
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		files:            nil,
		amount:           0,
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

// WithFiles add files to the builder
func (app *builder) WithFiles(files hashtree.HashTree) Builder {
	app.files = files
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

// Now builds a new Bucket instance
func (app *builder) Now() (Bucket, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build an Bucket instance")
	}

	if app.files == nil {
		return nil, errors.New("the files hashtree is mandatory in order to build an Bucket instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount must be greater than zero (0) in order to build an Bucket instance")
	}

	length := app.files.Length()
	if length < int(app.amount) {
		str := fmt.Sprintf("the chunk's length (%d) cannot be smaller than the amount (%d)", length, app.amount)
		return nil, errors.New(str)
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createBucket(immutable, app.files, app.amount), nil
}
