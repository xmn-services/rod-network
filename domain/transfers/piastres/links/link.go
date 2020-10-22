package links

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type link struct {
	immutable    entities.Immutable
	previousLink hash.Hash
	next         hash.Hash
	index        uint
}

func createLinkFromJSON(ins *jsonLink) (Link, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	previousLink, err := hashAdapter.FromString(ins.PreviousLink)
	if err != nil {
		return nil, err
	}

	next, err := hashAdapter.FromString(ins.Next)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithHash(*hsh).
		WithPreviousLink(*previousLink).
		WithNext(*next).
		WithIndex(ins.Index).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createLink(
	immutable entities.Immutable,
	previousLink hash.Hash,
	next hash.Hash,
	index uint,
) Link {
	out := link{
		immutable:    immutable,
		previousLink: previousLink,
		next:         next,
		index:        index,
	}

	return &out
}

// Hash returns the hash
func (obj *link) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// PreviousLink returns the previousLink hash
func (obj *link) PreviousLink() hash.Hash {
	return obj.previousLink
}

// Next returns the next hash
func (obj *link) Next() hash.Hash {
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
	ins := new(jsonLink)
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
	obj.previousLink = insLink.previousLink
	obj.next = insLink.next
	obj.index = insLink.index
	return nil
}
