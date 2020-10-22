package buckets

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	information      *hash.Hash
	absolutePath     string
	pk               encryption.PrivateKey
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		pk:               nil,
		information:      nil,
		absolutePath:     "",
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

// WithInformation adds an information to the builder
func (app *builder) WithInformation(information hash.Hash) Builder {
	app.information = &information
	return app
}

// WithAbsolutePath adds an absolutePath to the builder
func (app *builder) WithAbsolutePath(absolutePath string) Builder {
	app.absolutePath = absolutePath
	return app
}

// WithPrivateKey adds a privateKey to the builder
func (app *builder) WithPrivateKey(pk encryption.PrivateKey) Builder {
	app.pk = pk
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Bucket instance
func (app *builder) Now() (Bucket, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Bucket intance")
	}

	if app.pk == nil {
		return nil, errors.New("the PrivateKey is mandatory in order to build a Bucket instance")
	}

	if app.information == nil {
		return nil, errors.New("the information hash is mandatory in order to build a Bucket instance")
	}

	if app.absolutePath == "" {
		return nil, errors.New("the absolutePath is mandatory in order to build a Bucket instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createBucket(immutable, *app.information, app.absolutePath, app.pk), nil
}
