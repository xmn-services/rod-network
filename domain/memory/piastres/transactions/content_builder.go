package transactions

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/cancels"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
)

type contentBuilder struct {
	hashAdapter       hash.Adapter
	immutableBuilder  entities.ImmutableBuilder
	triggersOn        *time.Time
	executesOnTrigger bool
	fees              expenses.Expense
	expense           expenses.Expense
	cancel            cancels.Cancel
}

func createContentBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter:       hashAdapter,
		immutableBuilder:  immutableBuilder,
		triggersOn:        nil,
		executesOnTrigger: false,
		fees:              nil,
		expense:           nil,
		cancel:            nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter, app.immutableBuilder)
}

// TriggersOn adds a trigger time to the builder
func (app *contentBuilder) TriggersOn(triggersOn time.Time) ContentBuilder {
	app.triggersOn = &triggersOn
	return app
}

// ExecutesOnTrigger adds an executesOnTrigger flag to the builder
func (app *contentBuilder) ExecutesOnTrigger() ContentBuilder {
	app.executesOnTrigger = true
	return app
}

// WithFees add fees to the builder
func (app *contentBuilder) WithFees(fees expenses.Expense) ContentBuilder {
	app.fees = fees
	return app
}

// WithExpense adds an expense to the builder
func (app *contentBuilder) WithExpense(expense expenses.Expense) ContentBuilder {
	app.expense = expense
	return app
}

// WithCancel adds a cancel to the builder
func (app *contentBuilder) WithCancel(cancel cancels.Cancel) ContentBuilder {
	app.cancel = cancel
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.triggersOn == nil {
		return nil, errors.New("the triggersOn time is mandatory in order to build a Content instance")
	}

	executesOnTrigger := "false"
	if app.executesOnTrigger {
		executesOnTrigger = "true"
	}

	data := [][]byte{
		[]byte(executesOnTrigger),
		[]byte(strconv.Itoa(int(app.triggersOn.Nanosecond()))),
	}

	if app.fees != nil {
		data = append(data, app.fees.Hash().Bytes())
	}

	if app.expense != nil {
		data = append(data, app.expense.Hash().Bytes())
	}

	if app.cancel != nil {
		data = append(data, app.cancel.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.fees != nil {
		if app.expense != nil {
			return createContentWithExpenseAndFees(
				*hsh,
				*app.triggersOn,
				app.executesOnTrigger,
				app.expense,
				app.fees,
			), nil
		}

		if app.cancel != nil {
			return createContentWithCancelAndFees(
				*hsh,
				*app.triggersOn,
				app.executesOnTrigger,
				app.cancel,
				app.fees,
			), nil
		}
	}

	if app.expense != nil {
		return createContentWithExpense(
			*hsh,
			*app.triggersOn,
			app.executesOnTrigger,
			app.expense,
		), nil
	}

	if app.cancel != nil {
		return createContentWithCancel(
			*hsh,
			*app.triggersOn,
			app.executesOnTrigger,
			app.cancel,
		), nil
	}

	return nil, errors.New("the Content is invalid")
}
