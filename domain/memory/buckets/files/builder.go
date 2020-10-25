package files

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/buckets/files/chunks"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	relativePath     string
	chunks           []chunks.Chunk
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		relativePath:     "",
		chunks:           nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithRelativePath adds a relativePath to the builder
func (app *builder) WithRelativePath(relativePath string) Builder {
	app.relativePath = relativePath
	return app
}

// WithChunks add chunks to the builder
func (app *builder) WithChunks(chunks []chunks.Chunk) Builder {
	app.chunks = chunks
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new File instance
func (app *builder) Now() (File, error) {
	if app.relativePath == "" {
		return nil, errors.New("the relativePath is mandatory in order to build a File instance")
	}

	if app.chunks == nil {
		return nil, errors.New("the chunks are mandatory in order to build a File instance")
	}

	if len(app.chunks) <= 0 {
		return nil, errors.New("there must be at least 1 Chunk in order to build a File instance")
	}

	data := [][]byte{
		[]byte(app.relativePath),
	}

	for _, oneChunk := range app.chunks {
		data = append(data, oneChunk.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	mp := map[string]chunks.Chunk{}
	for _, oneChunk := range app.chunks {
		mp[oneChunk.Hash().String()] = oneChunk
	}

	return createFile(immutable, app.relativePath, app.chunks, mp), nil
}
