package expenses

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
	from             []hash.Hash
	to               *hash.Hash
	signatures       []signature.RingSignature
	remaining        *hash.Hash
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		from:             nil,
		to:               nil,
		signatures:       nil,
		remaining:        nil,
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

// From adds a from hash to the builder
func (app *builder) From(from []hash.Hash) Builder {
	app.from = from
	return app
}

// To adds a to hash to the builder
func (app *builder) To(to hash.Hash) Builder {
	app.to = &to
	return app
}

// WithSignatures adds signatures hash to the builder
func (app *builder) WithSignatures(signatures []signature.RingSignature) Builder {
	app.signatures = signatures
	return app
}

// WithRemaining adds a remaining hash to the builder
func (app *builder) WithRemaining(remaining hash.Hash) Builder {
	app.remaining = &remaining
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Expense instance
func (app *builder) Now() (Expense, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build an Expense instance")
	}

	if app.from == nil {
		return nil, errors.New("the from hashes are mandatory in order to build an Expense instance")
	}

	if app.to == nil {
		return nil, errors.New("the to hash is mandatory in order to build an Expense instance")
	}

	if len(app.from) <= 0 {
		return nil, errors.New("there must be at least 1 from hash in order to build an Expense instance")
	}

	if app.signatures == nil {
		return nil, errors.New("the ring signatures are mandatory in order to build an Expense instance")
	}

	if len(app.signatures) <= 0 {
		return nil, errors.New("there must be at least 1 ring signature in order to build an Expense instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	if app.remaining != nil {
		return createExpenseWithRemaining(immutable, app.from, *app.to, app.signatures, app.remaining), nil
	}

	return createExpense(immutable, app.from, *app.to, app.signatures), nil
}
