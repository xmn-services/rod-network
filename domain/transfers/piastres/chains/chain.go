package chains

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type chain struct {
	immutable entities.Immutable
	genesis   hash.Hash
	root      hash.Hash
	head      hash.Hash
	height    uint
}

func createChainFromJSON(ins *jsonChain) (Chain, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	genesis, err := hashAdapter.FromString(ins.Genesis)
	if err != nil {
		return nil, err
	}

	root, err := hashAdapter.FromString(ins.Root)
	if err != nil {
		return nil, err
	}

	head, err := hashAdapter.FromString(ins.Head)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithHash(*hsh).
		WithGenesis(*genesis).
		WithRoot(*root).
		WithHead(*head).
		WithHeight(ins.Height).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createChain(
	immutable entities.Immutable,
	genesis hash.Hash,
	root hash.Hash,
	head hash.Hash,
	height uint,
) Chain {
	out := chain{
		immutable: immutable,
		genesis:   genesis,
		root:      root,
		head:      head,
		height:    height,
	}

	return &out
}

// Hash returns the hash
func (obj *chain) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Genesis returns the genesis hash
func (obj *chain) Genesis() hash.Hash {
	return obj.genesis
}

// Root returns the root hash
func (obj *chain) Root() hash.Hash {
	return obj.root
}

// Head returns the head hash
func (obj *chain) Head() hash.Hash {
	return obj.head
}

// Height returns the height
func (obj *chain) Height() uint {
	return obj.height
}

// CreatedOn returns the creation time
func (obj *chain) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *chain) MarshalJSON() ([]byte, error) {
	ins := createJSONChainFromChain(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *chain) UnmarshalJSON(data []byte) error {
	ins := new(jsonChain)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createChainFromJSON(ins)
	if err != nil {
		return err
	}

	insChain := pr.(*chain)
	obj.immutable = insChain.immutable
	obj.genesis = insChain.genesis
	obj.root = insChain.root
	obj.head = insChain.head
	obj.height = insChain.height
	return nil
}
