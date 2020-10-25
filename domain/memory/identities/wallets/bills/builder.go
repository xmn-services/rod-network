package bills

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	bill             bills.Bill
	pks              []signature.PrivateKey
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		bill:             nil,
		pks:              nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithBill adds a bill to the builder
func (app *builder) WithBill(bill bills.Bill) Builder {
	app.bill = bill
	return app
}

// WithPrivateKeys add privateKeys to the builder
func (app *builder) WithPrivateKeys(pks []signature.PrivateKey) Builder {
	app.pks = pks
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Bill instance
func (app *builder) Now() (Bill, error) {
	if app.bill == nil {
		return nil, errors.New("the bill is mandatory in order to build a Bill instance")
	}

	if app.pks == nil {
		return nil, errors.New("the []PrivateKey are mandatory in order to build a Bill instance")
	}

	if len(app.pks) <= 0 {
		return nil, errors.New("there must be at least 1 PrivateKey in order to build a PrivateKey instance")
	}

	data := [][]byte{
		app.bill.Hash().Bytes(),
	}

	for _, onePK := range app.pks {
		data = append(data, []byte(onePK.String()))
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createBill(immutable, app.bill, app.pks), nil
}
