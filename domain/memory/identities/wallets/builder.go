package wallets

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/bills"
	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/statements"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	name             string
	bills            []bills.Bill
	statement        statements.Statement
	description      string
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		name:             "",
		bills:            nil,
		statement:        nil,
		description:      "",
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithBills add bills to the builder
func (app *builder) WithBills(bills []bills.Bill) Builder {
	app.bills = bills
	return app
}

// WithStatement adds a statement to the builder
func (app *builder) WithStatement(statement statements.Statement) Builder {
	app.statement = statement
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Wallet instance
func (app *builder) Now() (Wallet, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Wallet instance")
	}

	if app.statement != nil {
		return nil, errors.New("the statement is mandatory in order to build a Wallet instance")
	}

	if app.bills != nil && len(app.bills) <= 0 {
		app.bills = nil
	}

	data := [][]byte{
		[]byte(app.name),
		app.statement.Hash().Bytes(),
	}

	if app.bills != nil {
		for _, oneBill := range app.bills {
			data = append(data, []byte(oneBill.Hash().Bytes()))
		}
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	if app.bills != nil && app.description != "" {
		return createWalletWithBillsAndDescription(immutable, app.name, app.statement, app.bills, app.description), nil
	}

	if app.bills != nil {
		return createWalletWithBills(immutable, app.name, app.statement, app.bills), nil
	}

	if app.description != "" {
		return createWalletWithDescription(immutable, app.name, app.statement, app.description), nil
	}

	return createWallet(immutable, app.name, app.statement), nil
}
