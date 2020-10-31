package chains

import (
	"time"

	mined_block "github.com/xmn-services/rod-network/domain/memory/piastres/blocks/mined"
	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	mined_link "github.com/xmn-services/rod-network/domain/memory/piastres/links/mined"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type chain struct {
	immutable entities.Immutable
	genesis   genesis.Genesis
	root      mined_block.Block
	head      mined_link.Link
	height    uint
}

func createChain(
	immutable entities.Immutable,
	genesis genesis.Genesis,
	root mined_block.Block,
	head mined_link.Link,
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

// Genesis returns the genesis
func (obj *chain) Genesis() genesis.Genesis {
	return obj.genesis
}

// Root returns the root
func (obj *chain) Root() mined_block.Block {
	return obj.root
}

// Head returns the head
func (obj *chain) Head() mined_link.Link {
	return obj.head
}

// Height returns the height
func (obj *chain) Height() uint {
	return obj.Head().Link().Index()
}

// CreatedOn returns the creation time
func (obj *chain) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
