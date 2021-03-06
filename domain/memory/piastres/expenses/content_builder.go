package expenses

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type contentBuilder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	billBuilder      bills.Builder
	amount           uint64
	from             []bills.Bill
	lock             locks.Lock
	remaining        locks.Lock
	createdOn        *time.Time
}

func createContentBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
	billBuilder bills.Builder,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		billBuilder:      billBuilder,
		amount:           0,
		from:             nil,
		lock:             nil,
		remaining:        nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the contentBuilder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter, app.immutableBuilder, app.billBuilder)
}

// WithAmount adds an amount to the contentBuilder
func (app *contentBuilder) WithAmount(amount uint64) ContentBuilder {
	app.amount = amount
	return app
}

// From adds a from bill to the contentBuilder
func (app *contentBuilder) From(from []bills.Bill) ContentBuilder {
	app.from = from
	return app
}

// WithLock adds a new lock to the contentBuilder
func (app *contentBuilder) WithLock(lock locks.Lock) ContentBuilder {
	app.lock = lock
	return app
}

// WithRemaining adds a remaining lock to the contentBuilder
func (app *contentBuilder) WithRemaining(remaining locks.Lock) ContentBuilder {
	app.remaining = remaining
	return app
}

// CreatedOn adds a creation time to the contentBuilder
func (app *contentBuilder) CreatedOn(createdOn time.Time) ContentBuilder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Expense instance
func (app *contentBuilder) Now() (Content, error) {
	if app.from == nil {
		return nil, errors.New("the from bill is mandatory in order to build a Content instance")
	}

	if app.lock == nil {
		return nil, errors.New("the lock is mandatory in order to build a Content instance")
	}

	total := uint64(0)
	for _, oneBill := range app.from {
		total += oneBill.Amount()
	}

	if app.amount > total {
		str := fmt.Sprintf("the amount (%d) cannot be larger than the from amount (%d)", app.amount, total)
		return nil, errors.New(str)
	}

	remainingAmount := total - app.amount
	if app.remaining != nil {
		if remainingAmount <= 0 {
			return nil, errors.New("the remaining lock was expected to be nil since the bill was totally spent")
		}
	}

	if app.remaining == nil {
		if remainingAmount > 0 {
			str := fmt.Sprintf("the remaining lock was expected since there is a remaining amount (%d)", remainingAmount)
			return nil, errors.New(str)
		}
	}

	data := [][]byte{
		app.lock.Hash().Bytes(),
		[]byte(strconv.Itoa(int(app.amount))),
	}

	for _, oneBill := range app.from {
		data = append(data, oneBill.Hash().Bytes())
	}

	if app.remaining != nil {
		data = append(data, app.remaining.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	to, err := app.billBuilder.Create().WithAmount(app.amount).WithLock(app.lock).CreatedOn(immutable.CreatedOn()).Now()
	if err != nil {
		return nil, err
	}

	if app.remaining != nil {
		remaining, err := app.billBuilder.Create().WithAmount(remainingAmount).WithLock(app.remaining).CreatedOn(immutable.CreatedOn()).Now()
		if err != nil {
			return nil, err
		}

		return createContentWithRemaining(immutable, app.from, to, remaining), nil
	}

	return createContent(immutable, app.from, to), nil
}
