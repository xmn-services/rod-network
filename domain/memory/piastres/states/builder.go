package states

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	chain            *hash.Hash
	previous         *hash.Hash
	height           uint
	trx              []transactions.Transaction
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		chain:            nil,
		previous:         nil,
		height:           0,
		trx:              nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithChain adds a chain to the builder
func (app *builder) WithChain(chain hash.Hash) Builder {
	app.chain = &chain
	return app
}

// WithPrevious adds a previous hash to the builder
func (app *builder) WithPrevious(prev hash.Hash) Builder {
	app.previous = &prev
	return app
}

// WithHeight adds an height to the builder
func (app *builder) WithHeight(height uint) Builder {
	app.height = height
	return app
}

// WithTransactions add transactions to the builder
func (app *builder) WithTransactions(trx []transactions.Transaction) Builder {
	app.trx = trx
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new State instance
func (app *builder) Now() (State, error) {
	if app.chain == nil {
		return nil, errors.New("the chain is mandatory in order to build a State instance")
	}

	if app.previous == nil {
		return nil, errors.New("the previous hash is mandatory in order to build a State instance")
	}

	if app.height <= 0 {
		return nil, errors.New("the height msut be greater than zero (0) in order to build a State instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transactions are mandatory in order to build a State instance")
	}

	if len(app.trx) <= 0 {
		return nil, errors.New("there must be at least 1 Transaction instance in order to build a State instance")
	}

	data := [][]byte{
		app.chain.Bytes(),
		app.previous.Bytes(),
		[]byte(strconv.Itoa(int(app.height))),
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

	return createState(immutable, *app.chain, *app.previous, app.height, app.trx), nil
}
