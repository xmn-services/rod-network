package links

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
)

type link struct {
	immutable entities.Immutable
	prevLink  hash.Hash
	next      blocks.Block
	index     uint
}

func createLinkFromJSON(ins *JSONLink) (Link, error) {
	hashAdapter := hash.NewAdapter()
	prevLink, err := hashAdapter.FromString(ins.PreviousLink)
	if err != nil {
		return nil, err
	}

	blocksAdapter := blocks.NewAdapter()
	next, err := blocksAdapter.ToBlock(ins.Next)
	if err != nil {
		return nil, err
	}

	return NewBuilder().Create().
		WithIndex(ins.Index).
		WithNext(next).
		WithPreviousLink(*prevLink).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createLink(
	immutable entities.Immutable,
	prevLink hash.Hash,
	next blocks.Block,
	index uint,
) Link {
	out := link{
		immutable: immutable,
		prevLink:  prevLink,
		next:      next,
		index:     index,
	}

	return &out
}

// Hash returns the hash
func (obj *link) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// PreviousLink returns the previous link hash
func (obj *link) PreviousLink() hash.Hash {
	return obj.prevLink
}

// Next returns the next block
func (obj *link) Next() blocks.Block {
	return obj.next
}

// Index returns the index
func (obj *link) Index() uint {
	return obj.index
}

// CreatedOn returns the creation time
func (obj *link) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *link) MarshalJSON() ([]byte, error) {
	ins := createJSONLinkFromLink(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *link) UnmarshalJSON(data []byte) error {
	ins := new(JSONLink)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createLinkFromJSON(ins)
	if err != nil {
		return err
	}

	insLink := pr.(*link)
	obj.immutable = insLink.immutable
	obj.prevLink = insLink.prevLink
	obj.next = insLink.next
	obj.index = insLink.index
	return nil
}
