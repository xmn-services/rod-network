package bills

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	lock             *hash.Hash
	amount           uint
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		lock:             nil,
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

// WithLock adds a lock to the builder
func (app *builder) WithLock(lock hash.Hash) Builder {
	app.lock = &lock
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

// Now builds a new Bill instance
func (app *builder) Now() (Bill, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Bill instance")
	}

	if app.lock == nil {
		return nil, errors.New("the lock hash is mandatory in order to build a Bill instance")
	}

	if app.amount == 0 {
		return nil, errors.New("the amount is mandatory in order to build a Bill instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createBill(immutable, *app.lock, app.amount), nil
}
