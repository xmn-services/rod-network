package locks

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	pubKeys          []hash.Hash
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		pubKeys:          nil,
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

// WithPublicKeys adds pubkey's hashes to the builder
func (app *builder) WithPublicKeys(pubKeys []hash.Hash) Builder {
	app.pubKeys = pubKeys
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Lock instance
func (app *builder) Now() (Lock, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Lock instance")
	}

	if app.pubKeys == nil {
		return nil, errors.New("the pubkey hashes are mandatory in order to build a Lock instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createLock(immutable, app.pubKeys), nil
}
