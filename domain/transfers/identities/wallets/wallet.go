package wallets

import (
	"encoding/json"
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

func createWalletFromJSON(ins *jsonWallet) (Wallet, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	statement, err := hashAdapter.FromString(ins.Statement)
	if err != nil {
		return nil, err
	}

	builder := NewBuilder().Create().
		WithHash(*hsh).
		WithName(ins.Name).
		WithStatement(*statement).
		CreatedOn(ins.CreatedOn)

	if ins.Description != "" {
		builder.WithDescription(ins.Description)
	}

	if len(ins.Bills) > 0 {
		bills := []hash.Hash{}
		for _, oneBill := range ins.Bills {
			hsh, err := hashAdapter.FromString(oneBill)
			if err != nil {
				return nil, err
			}

			bills = append(bills, *hsh)
		}

		builder.WithBills(bills)
	}

	return builder.Now()
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

// MarshalJSON converts the instance to JSON
func (obj *wallet) MarshalJSON() ([]byte, error) {
	ins := createJSONWalletFromWallet(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *wallet) UnmarshalJSON(data []byte) error {
	ins := new(jsonWallet)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createWalletFromJSON(ins)
	if err != nil {
		return err
	}

	insWallet := pr.(*wallet)
	obj.immutable = insWallet.immutable
	obj.name = insWallet.name
	obj.bills = insWallet.bills
	obj.statement = insWallet.statement
	return nil
}
