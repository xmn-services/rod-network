package blocks

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	genesis          genesis.Genesis
	additional       uint
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
		genesis:          nil,
		additional:       0,
		trx:              nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithGenesis adds a genesis to the builder
func (app *builder) WithGenesis(gen genesis.Genesis) Builder {
	app.genesis = gen
	return app
}

// WithAdditional adds an additional amount of trx to the builder
func (app *builder) WithAdditional(additional uint) Builder {
	app.additional = additional
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

// Now builds a new Block instance
func (app *builder) Now() (Block, error) {
	if app.genesis == nil {
		return nil, errors.New("the genesis instance is mandatory in order to build a Block instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transactions are mandatory in order to build a Block instance")
	}

	if len(app.trx) <= 0 {
		return nil, errors.New("there must be at least 1 Tranaction instance in order to build a Block instance")
	}

	data := [][]byte{
		app.genesis.Hash().Bytes(),
		[]byte(strconv.Itoa(int(app.additional))),
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

	return createBlock(immutable, app.genesis, app.additional, app.trx), nil
}
