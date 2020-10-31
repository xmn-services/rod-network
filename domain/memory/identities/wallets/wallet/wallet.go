package wallet

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/wallet/bills"
	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/wallet/statements"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type wallet struct {
	immutable   entities.Immutable
	name        string
	bills       []bills.Bill
	statement   statements.Statement
	description string
}

func createWallet(
	immutable entities.Immutable,
	name string,
	statement statements.Statement,
) Wallet {
	return createWalletInternally(immutable, name, statement, nil, "")
}

func createWalletWithBills(
	immutable entities.Immutable,
	name string,
	statement statements.Statement,
	bills []bills.Bill,
) Wallet {
	return createWalletInternally(immutable, name, statement, bills, "")
}

func createWalletWithDescription(
	immutable entities.Immutable,
	name string,
	statement statements.Statement,
	description string,
) Wallet {
	return createWalletInternally(immutable, name, statement, nil, description)
}

func createWalletWithBillsAndDescription(
	immutable entities.Immutable,
	name string,
	statement statements.Statement,
	bills []bills.Bill,
	description string,
) Wallet {
	return createWalletInternally(immutable, name, statement, bills, description)
}

func createWalletInternally(
	immutable entities.Immutable,
	name string,
	statement statements.Statement,
	bills []bills.Bill,
	description string,
) Wallet {
	out := wallet{
		immutable:   immutable,
		name:        name,
		bills:       bills,
		statement:   statement,
		description: description,
	}

	return &out
}

// Hash returns the hash
func (obj *wallet) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Name returns the name
func (obj *wallet) Name() string {
	return obj.name
}

// Statement returns the statement
func (obj *wallet) Statement() statements.Statement {
	return obj.statement
}

// CreatedOn returns the creation time
func (obj *wallet) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasBills returns true if there is bills, false otherwise
func (obj *wallet) HasBills() bool {
	return obj.bills != nil
}

// Bills returns the bills, if any
func (obj *wallet) Bills() []bills.Bill {
	return obj.bills
}

// HasDescription returns true if there is a description, false otherwise
func (obj *wallet) HasDescription() bool {
	return obj.description != ""
}

// Description returns the description, if any
func (obj *wallet) Description() string {
	return obj.description
}
