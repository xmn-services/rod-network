package transactions

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	fees             []hash.Hash
	bucket           *hash.Hash
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		fees:             nil,
		bucket:           nil,
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

// WithFees adds fees to the builder
func (app *builder) WithFees(fees []hash.Hash) Builder {
	app.fees = fees
	return app
}

// WithBucket adds a bucket to the builder
func (app *builder) WithBucket(bucket hash.Hash) Builder {
	app.bucket = &bucket
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Transaction instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	if app.fees != nil {
		if app.bucket != nil {
			return createTransactionWithBucketAndFees(
				immutable,
				app.bucket,
				app.fees,
			), nil
		}

		return createTransactionWithFees(
			immutable,
			app.fees,
		), nil
	}

	if app.bucket != nil {
		return createTransactionWithBucket(
			immutable,
			app.bucket,
		), nil
	}

	return nil, errors.New("the Transaction instance is invalid")
}
