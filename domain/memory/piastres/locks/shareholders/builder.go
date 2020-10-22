package shareholders

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	key              *hash.Hash
	power            uint
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		key:              nil,
		power:            0,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.immutableBuilder,
	)
}

// WithKey adds key to the builder
func (app *builder) WithKey(key hash.Hash) Builder {
	app.key = &key
	return app
}

// WithPower adds power to the builder
func (app *builder) WithPower(power uint) Builder {
	app.power = power
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new ShareHolder instance
func (app *builder) Now() (ShareHolder, error) {
	if app.key == nil {
		return nil, errors.New("the publicKey is mandatory in order to build a ShareHolder instance")
	}

	if app.power <= 0 {
		return nil, errors.New("the power must be greater than zero (0) in order to build a ShareHolder instance")
	}

	data := [][]byte{
		[]byte(strconv.Itoa(int(app.power))),
		app.key.Bytes(),
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createShareHolder(immutable, *app.key, app.power), nil
}
