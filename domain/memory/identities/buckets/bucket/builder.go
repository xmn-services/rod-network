package bucket

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/buckets"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	pkAdapter        encryption.Adapter
	immutableBuilder entities.ImmutableBuilder
	bucket           buckets.Bucket
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
		bucket:           nil,
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

// WithBucket adds a bucket to the builder
func (app *builder) WithBucket(bucket buckets.Bucket) Builder {
	app.bucket = bucket
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
	if app.bucket == nil {
		return nil, errors.New("the bucket is mandatory in order to build a Bucket instance")
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

	return createBucket(immutable, app.bucket, absolutePath, app.pk), nil
}
