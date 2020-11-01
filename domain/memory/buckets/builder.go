package buckets

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/buckets/files"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	files            []files.File
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		files:            nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithFiles add files to the builder
func (app *builder) WithFiles(files []files.File) Builder {
	app.files = files
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Bucket instance
func (app *builder) Now() (Bucket, error) {
	if app.files == nil {
		return nil, errors.New("the files are mandatory in order to build an Bucket instance")
	}

	if len(app.files) <= 0 {
		return nil, errors.New("there must be at least 1 File in order to build an Bucket instance")
	}

	data := [][]byte{}
	for _, oneFile := range app.files {
		data = append(data, oneFile.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	mp := map[string]files.File{}
	for _, oneFile := range app.files {
		mp[oneFile.RelativePath()] = oneFile
	}

	return createBucket(immutable, app.files, mp), nil
}
