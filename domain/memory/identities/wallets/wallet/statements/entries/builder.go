package entries

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	name             string
	trx              []transactions.Transaction
	description      string
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		name:             "",
		trx:              nil,
		description:      "",
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// WithTransactions add transactions to the builder
func (app *builder) WithTransactions(trx []transactions.Transaction) Builder {
	app.trx = trx
	return app
}

// CreatedOn adds a cration time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Entry instance
func (app *builder) Now() (Entry, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Entry instance")
	}

	if app.trx == nil {
		return nil, errors.New("the []Transaction are mandatory in order to build an Entry instance")
	}

	if len(app.trx) <= 0 {
		return nil, errors.New("there must be at least 1 Transaction in order to build an Entry instance")
	}

	data := [][]byte{
		[]byte(app.name),
	}

	for _, oneTrx := range app.trx {
		data = append(data, oneTrx.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	if app.description != "" {
		return createEntryWithDescription(immutable, app.name, app.trx, app.description), nil
	}

	return createEntry(immutable, app.name, app.trx), nil
}
