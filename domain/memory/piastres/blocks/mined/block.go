package mined

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type block struct {
	immutable entities.Immutable
	block     blocks.Block
	mining    string
}

func createBlock(
	immutable entities.Immutable,
	blk blocks.Block,
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

// Block returns the block
func (obj *block) Block() blocks.Block {
	return obj.block
}

// Mining returns the mining
func (obj *block) Mining() string {
	return obj.mining
}

// CreatedOn returns the creation time
func (obj *block) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
