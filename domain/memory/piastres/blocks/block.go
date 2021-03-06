package blocks

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type block struct {
	immutable  entities.Immutable
	genesis    genesis.Genesis
	additional uint
	trx        []transactions.Transaction
}

func createBlockFromJSON(ins *JSONBlock) (Block, error) {
	trxAdapter := transactions.NewAdapter()
	trx := []transactions.Transaction{}
	for _, oneJSTrx := range ins.Trx {
		oneTrx, err := trxAdapter.ToTransaction(oneJSTrx)
		if err != nil {
			return nil, err
		}

		trx = append(trx, oneTrx)
	}

	genAdapter := genesis.NewAdapter()
	gen, err := genAdapter.ToGenesis(ins.Genesis)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithTransactions(trx).
		WithGenesis(gen).
		WithAdditional(ins.Additional).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createBlock(
	immutable entities.Immutable,
	genesis genesis.Genesis,
	additional uint,
	trx []transactions.Transaction,
) Block {
	out := block{
		immutable:  immutable,
		genesis:    genesis,
		additional: additional,
		trx:        trx,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Genesis returns the genesis
func (obj *block) Genesis() genesis.Genesis {
	return obj.genesis
}

// Additional returns the additional trx
func (obj *block) Additional() uint {
	return obj.additional
}

// Transactions returns the transactions
func (obj *block) Transactions() []transactions.Transaction {
	return obj.trx
}

// CreatedOn returns the creation time
func (obj *block) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *block) MarshalJSON() ([]byte, error) {
	ins := createJSONBlockFromBlock(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *block) UnmarshalJSON(data []byte) error {
	ins := new(JSONBlock)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createBlockFromJSON(ins)
	if err != nil {
		return err
	}

	insBlock := pr.(*block)
	obj.immutable = insBlock.immutable
	obj.genesis = insBlock.genesis
	obj.additional = insBlock.additional
	obj.trx = insBlock.trx
	return nil
}
