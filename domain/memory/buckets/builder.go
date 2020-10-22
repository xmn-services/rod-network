package buckets

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/buckets/informations"
)

type builder struct {
	hashAdapter      hash.Adapter
	pkAdapter        encryption.Adapter
	immutableBuilder entities.ImmutableBuilder
	information      informations.Information
	absolutePath     string
	pk               encryption.PrivateKey
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	pkAdapter encryption.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		pkAdapter:        pkAdapter,
		immutableBuilder: immutableBuilder,
		information:      nil,
		absolutePath:     "",
		pk:               nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.pkAdapter, app.immutableBuilder)
}

// WithInformation adds an information to the builder
func (app *builder) WithInformation(information informations.Information) Builder {
	app.information = information
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
	if app.information == nil {
		return nil, errors.New("the information is mandatory in order to build a Bucket instance")
	}

	if app.absolutePath == "" {
		return nil, errors.New("the absolutePath is mandatory in order to build a Bucket instance")
	}

	if app.pk == nil {
		return nil, errors.New("the PrivateKey is mandatory in order to build a Bucket instance")
	}

	absolutePath, err := filepath.Abs(app.absolutePath)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		str := fmt.Sprintf("the absolutePath (%s) does not exists", absolutePath)
		return nil, errors.New(str)
	}

	hsh, err := app.hashAdapter.FromBytes([]byte(absolutePath))
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createBucket(immutable, app.information, absolutePath, app.pk), nil
}
