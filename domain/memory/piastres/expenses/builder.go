package expenses

import (
	"errors"
	"fmt"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	content          Content
	signatures       [][]signature.RingSignature
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		content:          nil,
		signatures:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithContent adds content to the builder
func (app *builder) WithContent(content Content) Builder {
	app.content = content
	return app
}

// WithSignatures add ring signatures to the builder
func (app *builder) WithSignatures(sigs [][]signature.RingSignature) Builder {
	app.signatures = sigs
	return app
}

// Now builds a new Expense instance
func (app *builder) Now() (Expense, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build an Expense instance")
	}

	if app.signatures == nil {
		return nil, errors.New("the ring signatures are mandatory in order to build an Expense instance")
	}

	from := app.content.From()
	if len(app.signatures) != len(from) {
		str := fmt.Sprintf("there must be the same amount of ring signatures (%d) as from bills (%d)", len(app.signatures), len(from))
		return nil, errors.New(str)
	}

	for index, oneBill := range from {
		err := oneBill.Lock().Unlock(app.signatures[index])
		if err != nil {
			return nil, err
		}
	}

	data := [][]byte{
		app.content.Hash().Bytes(),
	}

	for _, oneSignatureList := range app.signatures {
		for _, oneSignature := range oneSignatureList {
			data = append(data, []byte(oneSignature.String()))
		}
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	createdOn := app.content.CreatedOn()
	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(&createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createExpense(immutable, app.content, app.signatures), nil
}
