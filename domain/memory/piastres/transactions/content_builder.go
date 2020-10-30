package transactions

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/libs/hash"
)

type contentBuilder struct {
	hashAdapter hash.Adapter
	triggersOn  *time.Time
	element     Element
	fees        []expenses.Expense
}

func createContentBuilder(
	hashAdapter hash.Adapter,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter: hashAdapter,
		triggersOn:  nil,
		element:     nil,
		fees:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter)
}

// TriggersOn adds a triggersOn to the builder
func (app *contentBuilder) TriggersOn(triggersOn time.Time) ContentBuilder {
	app.triggersOn = &triggersOn
	return app
}

// WithElement adds an element to the builder
func (app *contentBuilder) WithElement(element Element) ContentBuilder {
	app.element = element
	return app
}

// WithFees adds fees to the builder
func (app *contentBuilder) WithFees(fees []expenses.Expense) ContentBuilder {
	app.fees = fees
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.triggersOn == nil {
		return nil, errors.New("the triggersOn is mandatory in order to build a Content instance")
	}

	if app.fees != nil && len(app.fees) <= 0 {
		app.fees = nil
	}

	data := [][]byte{
		[]byte(strconv.Itoa(int(app.triggersOn.UnixNano()))),
	}

	if app.element != nil {
		data = append(data, app.element.Hash().Bytes())
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

	if app.fees != nil && app.element != nil {
		return createContentWithElementAndFees(*hsh, *app.triggersOn, app.element, app.fees), nil
	}

	if app.fees != nil {
		return createContentWithFees(*hsh, *app.triggersOn, app.fees), nil
	}

	if app.element != nil {
		return createContentWithElement(*hsh, *app.triggersOn, app.element), nil
	}

	return nil, errors.New("the Content is invalid")
}
