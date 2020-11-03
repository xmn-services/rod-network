package chains

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	mined_block "github.com/xmn-services/rod-network/domain/memory/piastres/blocks/mined"
	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	mined_link "github.com/xmn-services/rod-network/domain/memory/piastres/links/mined"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	genesis          genesis.Genesis
	root             mined_block.Block
	head             mined_link.Link
	total            uint
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
		total:            0,
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

// WithTotal adds a total to block instance
func (app *builder) WithTotal(total uint) Builder {
	app.total = total
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

	known := uint(len(app.root.Block().Transactions()) + len(app.head.Link().Next().Transactions()))
	if app.total < known {
		str := fmt.Sprintf("there is %d transactions in the root and head blocks, therefore the total (%d) must be bigger or equal to that amount", known, app.total)
		return nil, errors.New(str)
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.genesis.Hash().Bytes(),
		app.root.Hash().Bytes(),
		app.head.Hash().Bytes(),
		[]byte(strconv.Itoa(int(app.total))),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createChain(immutable, app.genesis, app.root, app.head, app.total), nil
}
