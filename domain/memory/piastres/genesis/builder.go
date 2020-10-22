package genesis

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
)

type builder struct {
	hashAdapter             hash.Adapter
	immutableBuilder        entities.ImmutableBuilder
	blockDiffBase           uint
	blockDiffIncreasePerTrx float64
	linkDiff                uint
	bill                    bills.Bill
	createdOn               *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:             hashAdapter,
		immutableBuilder:        immutableBuilder,
		blockDiffBase:           0,
		blockDiffIncreasePerTrx: 0.0,
		linkDiff:                0,
		bill:                    nil,
		createdOn:               nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithBlockDifficultyBase adds a block difficulty base to the builder
func (app *builder) WithBlockDifficultyBase(blockDiffBase uint) Builder {
	app.blockDiffBase = blockDiffBase
	return app
}

// WithBlockDifficultyIncreasePerTrx adds a block difficulty increasePerTrx to the builder
func (app *builder) WithBlockDifficultyIncreasePerTrx(blockDiffIncreasePerTrx float64) Builder {
	app.blockDiffIncreasePerTrx = blockDiffIncreasePerTrx
	return app
}

// WithLinkDifficulty adds a link difficulty to the builder
func (app *builder) WithLinkDifficulty(linkDiff uint) Builder {
	app.linkDiff = linkDiff
	return app
}

// WithBill adds a bill to the builder
func (app *builder) WithBill(bill bills.Bill) Builder {
	app.bill = bill
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Genesis instance
func (app *builder) Now() (Genesis, error) {
	if app.blockDiffBase <= 0 {
		return nil, errors.New("the block difficulty base must be greater than zero (0) in order to build a Genesis instance")
	}

	if app.blockDiffIncreasePerTrx <= 0.0 {
		return nil, errors.New("the block difficulty increasePerTrx must be greater than zero (0.0) in order to build a Genesis instance")
	}

	if app.linkDiff <= 0 {
		return nil, errors.New("the link difficulty must be greater than zero (0.0) in order to build a Genesis instance")
	}

	if app.bill == nil {
		return nil, errors.New("the bill is mandatory in order to build a Genesis instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(app.blockDiffBase))),
		[]byte(strconv.FormatFloat(app.blockDiffIncreasePerTrx, 'f', 12, 64)),
		[]byte(strconv.Itoa(int(app.linkDiff))),
		app.bill.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	block := createBlock(app.blockDiffBase, app.blockDiffIncreasePerTrx)
	diff := createDifficulty(block, app.linkDiff)
	return createGenesis(immutable, app.bill, diff), nil
}
