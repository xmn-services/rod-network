package cancels

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	expense          expenses.Expense
	lock             locks.Lock
	signatures       []signature.RingSignature
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		expense:          nil,
		lock:             nil,
		signatures:       nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithExpense adds an expense to the builder
func (app *builder) WithExpense(expense expenses.Expense) Builder {
	app.expense = expense
	return app
}

// WithLock adds a lock to the builder
func (app *builder) WithLock(lock locks.Lock) Builder {
	app.lock = lock
	return app
}

// WithSignatures add ring signatures to the builder
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
	if app.expense == nil {
		return nil, errors.New("the expense is mandatory in order to build a Cancel instancfe")
	}

	if app.lock == nil {
		return nil, errors.New("the lock is mandatory in order to build a Cancel instance")
	}

	if app.signatures == nil {
		return nil, errors.New("the ring signatures are mandatory in order to build a Cancel instance")
	}

	err := app.expense.Content().Cancel().Unlock(app.signatures)
	if err != nil {
		return nil, err
	}

	data := [][]byte{
		app.expense.Hash().Bytes(),
		app.lock.Hash().Bytes(),
	}

	for _, oneSignature := range app.signatures {
		data = append(data, []byte(oneSignature.String()))
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createCancel(immutable, app.expense, app.lock, app.signatures), nil
}
