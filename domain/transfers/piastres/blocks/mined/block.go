package mined

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type block struct {
	immutable entities.Immutable
	block     hash.Hash
	mining    string
}

func createBlockFromJSON(ins *jsonBlock) (Block, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	block, err := hashAdapter.FromString(ins.Block)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithHash(*hsh).
		WithBlock(*block).
		WithMining(ins.Mining).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createBlock(
	immutable entities.Immutable,
	blk hash.Hash,
	mining string,
) Block {
	out := block{
		immutable: immutable,
		block:     blk,
		mining:    mining,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Block returns the block hash
func (obj *block) Block() hash.Hash {
	return obj.block
}

// Mining returns the mining results
func (obj *block) Mining() string {
	return obj.mining
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
	ins := new(jsonBlock)
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
	obj.block = insBlock.block
	obj.mining = insBlock.mining
	return nil
}
