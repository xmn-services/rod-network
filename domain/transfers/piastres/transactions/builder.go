package transactions

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	signature        signature.RingSignature
	triggersOn       *time.Time
	fees             []hash.Hash
	bucket           *hash.Hash
	cancel           *hash.Hash
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		signature:        nil,
		triggersOn:       nil,
		fees:             nil,
		bucket:           nil,
		cancel:           nil,
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

// TriggersOn adds a triggersOn time to the builder
func (app *builder) TriggersOn(triggersOn time.Time) Builder {
	app.triggersOn = &triggersOn
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

// WithCancel adds a cancel to the builder
func (app *builder) WithCancel(cancel hash.Hash) Builder {
	app.cancel = &cancel
	return app
}

// WithSignature adds a signature to the builder
func (app *builder) WithSignature(signature signature.RingSignature) Builder {
	app.signature = signature
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

	if app.signature == nil {
		return nil, errors.New("the signature is mandatory in order to build a Transaction instance")
	}

	if app.triggersOn == nil {
		return nil, errors.New("the triggersOn is mandatory in order to build a Transaction instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	if app.fees != nil {
		if app.bucket != nil {
			return createTransactionWithBucketAndFees(
				immutable,
				app.signature,
				*app.triggersOn,
				app.bucket,
				app.fees,
			), nil
		}

		if app.cancel != nil {
			return createTransactionWithCancelAndFees(
				immutable,
				app.signature,
				*app.triggersOn,
				app.cancel,
				app.fees,
			), nil
		}

		return createTransactionWithFees(
			immutable,
			app.signature,
			*app.triggersOn,
			app.fees,
		), nil
	}

	if app.bucket != nil {
		return createTransactionWithBucket(
			immutable,
			app.signature,
			*app.triggersOn,
			app.bucket,
		), nil
	}

	if app.cancel != nil {
		return createTransactionWithCancel(
			immutable,
			app.signature,
			*app.triggersOn,
			app.cancel,
		), nil
	}

	return nil, errors.New("the Transaction instance is invalid")
}
