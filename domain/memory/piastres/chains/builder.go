package chains

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	mined_block "github.com/xmn-services/rod-network/domain/memory/piastres/blocks/mined"
	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	mined_link "github.com/xmn-services/rod-network/domain/memory/piastres/links/mined"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	genesis          genesis.Genesis
	root             mined_block.Block
	head             mined_link.Link
	height           uint
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		genesis:          nil,
		root:             nil,
		head:             nil,
		height:           0,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithGenesis adds a genesis instance
func (app *builder) WithGenesis(gen genesis.Genesis) Builder {
	app.genesis = gen
	return app
}

// WithRoot adds a root block instance
func (app *builder) WithRoot(root mined_block.Block) Builder {
	app.root = root
	return app
}

// WithHead adds an head block instance
func (app *builder) WithHead(head mined_link.Link) Builder {
	app.head = head
	return app
}

// WithHeight adds an height block instance
func (app *builder) WithHeight(height uint) Builder {
	app.height = height
	return app
}

// CreatedOn adds a creation time instance
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Chain instance
func (app *builder) Now() (Chain, error) {
	if app.genesis == nil {
		return nil, errors.New("the genesis instance is mandatory in order to build a Chain instance")
	}

	if app.root == nil {
		return nil, errors.New("the root block is mandatory in order to build a Chain instance")
	}

	if app.head == nil {
		return nil, errors.New("the head link is mandatory in order to build a Chain instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.genesis.Hash().Bytes(),
		app.root.Hash().Bytes(),
		app.head.Hash().Bytes(),
		[]byte(strconv.Itoa(int(app.height))),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createChain(immutable, app.genesis, app.root, app.head, app.height), nil
}
