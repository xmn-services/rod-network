package mined

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/links"
)

type link struct {
	immutable entities.Immutable
	link      links.Link
	mining    string
}

func createLinkFromJSON(ins *JSONLink) (Link, error) {
	linksAdapter := links.NewAdapter()
	link, err := linksAdapter.ToLink(ins.Link)
	if err != nil {
		return nil, err
	}

	return NewBuilder().Create().
		WithLink(link).
		WithMining(ins.Mining).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createLink(
	immutable entities.Immutable,
	lnk links.Link,
	mining string,
) Link {
	out := link{
		immutable: immutable,
		link:      lnk,
		mining:    mining,
	}

	return &out
}

// Hash returns the hash
func (obj *link) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Link returns the link instance
func (obj *link) Link() links.Link {
	return obj.link
}

// Mining returns the mining results
func (obj *link) Mining() string {
	return obj.mining
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
	obj.link = insLink.link
	obj.mining = insLink.mining
	return nil
}
