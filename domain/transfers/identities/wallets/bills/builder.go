package bills

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
	bill             *hash.Hash
	pks              []signature.PrivateKey
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		bill:             nil,
		pks:              nil,
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

// WithBill adds a bill to the builder
func (app *builder) WithBill(bill hash.Hash) Builder {
	app.bill = &bill
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
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Bill instance")
	}

	if app.bill == nil {
		return nil, errors.New("the bill is mandatory in order to build a Bill instance")
	}

	if app.pks == nil {
		return nil, errors.New("the []PrivateKey are mandatory in order to build a Bill instance")
	}

	if len(app.pks) <= 0 {
		return nil, errors.New("there must be at least 1 PrivateKey in order to build a PrivateKey instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createBill(immutable, *app.bill, app.pks), nil
}
