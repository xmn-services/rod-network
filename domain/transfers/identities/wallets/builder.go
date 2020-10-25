package wallets

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	name             string
	bills            []hash.Hash
	statement        *hash.Hash
	description      string
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
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
	return createBuilder(app.immutableBuilder)
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.hash = &hash
	return app
}

// WithBills add bills to the builder
func (app *builder) WithBills(bills []hash.Hash) Builder {
	app.bills = bills
	return app
}

// WithStatement adds a statement to the builder
func (app *builder) WithStatement(statement hash.Hash) Builder {
	app.statement = &statement
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
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Wallet instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Wallet instance")
	}

	if app.statement != nil {
		return nil, errors.New("the statement is mandatory in order to build a Wallet instance")
	}

	if app.bills != nil && len(app.bills) <= 0 {
		app.bills = nil
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	if app.bills != nil && app.description != "" {
		return createWalletWithBillsAndDescription(immutable, app.name, *app.statement, app.bills, app.description), nil
	}

	if app.bills != nil {
		return createWalletWithBills(immutable, app.name, *app.statement, app.bills), nil
	}

	if app.description != "" {
		return createWalletWithDescription(immutable, app.name, *app.statement, app.description), nil
	}

	return createWallet(immutable, app.name, *app.statement), nil
}
