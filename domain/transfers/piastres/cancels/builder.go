package cancels

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
	expense          *hash.Hash
	lock             *hash.Hash
	signatures       []signature.RingSignature
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		expense:          nil,
		lock:             nil,
		signatures:       nil,
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

// WithExpense adds an expense hash to the builder
func (app *builder) WithExpense(expense hash.Hash) Builder {
	app.expense = &expense
	return app
}

// WithLock adds a lock hash to the builder
func (app *builder) WithLock(lock hash.Hash) Builder {
	app.lock = &lock
	return app
}

// WithSignatures adds signatures hash to the builder
func (app *builder) WithSignatures(signatures []signature.RingSignature) Builder {
	app.signatures = signatures
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Cancel instance
func (app *builder) Now() (Cancel, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Cancel instance")
	}

	if app.expense == nil {
		return nil, errors.New("the expense hash is mandatory in order to build a Cancel instance")
	}

	if app.lock == nil {
		return nil, errors.New("the lock hash is mandatory in order to build a Cancel instance")
	}

	if app.signatures == nil {
		return nil, errors.New("the ring signatures are mandatory in order to build a Cancel instance")
	}

	if len(app.signatures) <= 0 {
		return nil, errors.New("there must be at least 1 ring signature in order to build a Cancel instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createCancel(immutable, *app.expense, *app.lock, app.signatures), nil

}
