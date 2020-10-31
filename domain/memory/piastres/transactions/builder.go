package transactions

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	triggersOn       *time.Time
	bucket           *hash.Hash
	fees             []expenses.Expense
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		triggersOn:       nil,
		bucket:           nil,
		fees:             nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithBucket adds a bucket to the builder
func (app *builder) WithBucket(bucket hash.Hash) Builder {
	app.bucket = &bucket
	return app
}

// WithFees adds fees to the builder
func (app *builder) WithFees(fees []expenses.Expense) Builder {
	app.fees = fees
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.fees != nil && len(app.fees) <= 0 {
		app.fees = nil
	}

	data := [][]byte{}
	if app.bucket != nil {
		data = append(data, app.bucket.Bytes())
	}

	if app.fees != nil {
		for _, oneFee := range app.fees {
			data = append(data, oneFee.Hash().Bytes())
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

	if app.fees != nil && app.bucket != nil {
		return createTransactionWithBucketAndFees(immutable, app.bucket, app.fees), nil
	}

	if app.fees != nil {
		return createTransactionWithFees(immutable, app.fees), nil
	}

	if app.bucket != nil {
		return createTransactionWithBucket(immutable, app.bucket), nil
	}

	return nil, errors.New("the Transaction is invalid")
}
