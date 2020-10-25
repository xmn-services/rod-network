package wallets

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type wallet struct {
	immutable   entities.Immutable
	name        string
	bills       []hash.Hash
	statement   hash.Hash
	description string
}

func createWallet(
	immutable entities.Immutable,
	name string,
	statement hash.Hash,
) Wallet {
	return createWalletInternally(immutable, name, statement, nil, "")
}

func createWalletWithBills(
	immutable entities.Immutable,
	name string,
	statement hash.Hash,
	bills []hash.Hash,
) Wallet {
	return createWalletInternally(immutable, name, statement, bills, "")
}

func createWalletWithDescription(
	immutable entities.Immutable,
	name string,
	statement hash.Hash,
	description string,
) Wallet {
	return createWalletInternally(immutable, name, statement, nil, description)
}

func createWalletWithBillsAndDescription(
	immutable entities.Immutable,
	name string,
	statement hash.Hash,
	bills []hash.Hash,
	description string,
) Wallet {
	return createWalletInternally(immutable, name, statement, bills, description)
}

func createWalletInternally(
	immutable entities.Immutable,
	name string,
	statement hash.Hash,
	bills []hash.Hash,
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
func (obj *wallet) Statement() hash.Hash {
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
func (obj *wallet) Bills() []hash.Hash {
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
