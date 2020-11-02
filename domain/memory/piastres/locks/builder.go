package locks

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	pubKeys          []hash.Hash
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		pubKeys:          nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithPublicKeys add publicKeys to the builder
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
	if app.pubKeys == nil {
		return nil, errors.New("the []PublicKey are mandatory in order to build a Lock instance")
	}

	if len(app.pubKeys) <= 0 {
		return nil, errors.New("there must be at least 1 PublicKey in order to build a Lock instance")
	}

	data := [][]byte{}
	for _, onePubKey := range app.pubKeys {
		data = append(data, onePubKey.Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	mpKeys := map[string]hash.Hash{}
	for _, onePublicKey := range app.pubKeys {
		mpKeys[onePublicKey.String()] = onePublicKey
	}

	if len(mpKeys) != len(app.pubKeys) {
		return nil, errors.New("at least 1 PublicKey contained in the Lock is duplicate")
	}

	return createLock(immutable, app.pubKeys, mpKeys), nil
}
