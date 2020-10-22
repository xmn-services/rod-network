package locks

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	shareHolders     []shareholders.ShareHolder
	treeshold        uint
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		shareHolders:     nil,
		treeshold:        0,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithShareHolders add shareholders to the builder
func (app *builder) WithShareHolders(shareHolders []shareholders.ShareHolder) Builder {
	app.shareHolders = shareHolders
	return app
}

// WithTreeshold adds a treeshold to the builder
func (app *builder) WithTreeshold(treeshold uint) Builder {
	app.treeshold = treeshold
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Lock instance
func (app *builder) Now() (Lock, error) {
	if app.shareHolders == nil {
		return nil, errors.New("the []ShareHolder are mandatory in order to build a Lock instance")
	}

	if len(app.shareHolders) <= 0 {
		return nil, errors.New("there must be at least 1 ShareHolder in order to build a Lock instance")
	}

	if app.treeshold <= 0 {
		return nil, errors.New("the treeshold must be greater than zero (0) in order to build a Lock instance")
	}

	total := uint(0)
	for _, oneHolder := range app.shareHolders {
		total += oneHolder.Power()
	}

	if app.treeshold > total {
		str := fmt.Sprintf("the treeshold (%d) cannot be bigger than the total amount of shares (%d)", app.treeshold, total)
		return nil, errors.New(str)
	}

	data := [][]byte{
		[]byte(strconv.Itoa(int(app.treeshold))),
	}

	for _, oneShareHolder := range app.shareHolders {
		data = append(data, oneShareHolder.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	holders := map[string]shareholders.ShareHolder{}
	for _, oneShareHolder := range app.shareHolders {
		keyname := oneShareHolder.Key().String()
		holders[keyname] = oneShareHolder
	}

	return createLock(immutable, app.shareHolders, holders, app.treeshold), nil
}
