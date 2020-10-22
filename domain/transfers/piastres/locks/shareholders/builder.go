package shareholders

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	key              *hash.Hash
	power            uint
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		key:              nil,
		power:            0,
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

// WithKey adds a key hash to the builder
func (app *builder) WithKey(key hash.Hash) Builder {
	app.key = &key
	return app
}

// WithPower adds a power to the builder
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
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a ShareHolder instance")
	}

	if app.key == nil {
		return nil, errors.New("the key hash is mandatory in order to build a ShareHolder instance")
	}

	if app.power <= 0 {
		return nil, errors.New("the power must be greater than zero (0) in order to build a ShareHolder instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createShareHolder(immutable, *app.key, app.power), nil
}
