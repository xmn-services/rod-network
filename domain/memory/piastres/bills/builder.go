package bills

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	lock             locks.Lock
	amount           uint
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		lock:             nil,
		amount:           0,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithLock adds a lock to the builder
func (app *builder) WithLock(lock locks.Lock) Builder {
	app.lock = lock
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
	if app.lock == nil {
		return nil, errors.New("the lock is mandatory in order to build a Bill instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount must be greater than zero (0) in order to build a Bill instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.lock.Hash().Bytes(),
		[]byte(strconv.Itoa(int(app.amount))),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createBill(immutable, app.lock, app.amount), nil
}
