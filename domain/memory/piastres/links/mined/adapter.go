package mined

import (
	"encoding/base64"
	"encoding/json"

	transfer_mined_link "github.com/xmn-services/rod-network/domain/transfers/piastres/links/mined"
)

type adapter struct {
	trBuilder transfer_mined_link.Builder
}

func createAdapter(
	trBuilder transfer_mined_link.Builder,
) Adapter {
	out := adapter{
		trBuilder: trBuilder,
	}

	return &out
}

// ToTransfer converts a link instance to a transfer link instance
func (app *adapter) ToTransfer(link Link) (transfer_mined_link.Link, error) {
	hsh := link.Hash()
	linkHash := link.Link().Hash()
	mining := link.Mining()
	createdOn := link.CreatedOn()
	return app.trBuilder.Create().WithHash(hsh).WithLink(linkHash).WithMining(mining).CreatedOn(createdOn).Now()
}

// Decode converts encoded data to a Link instance
func (app *adapter) Decode(encoded string) (Link, error) {
	js, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	ins := new(link)
	json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// Encode encodes a link instance
func (app *adapter) Encode(link Link) (string, error) {
	js, err := json.Marshal(link)
	if err != nil {
		return "", nil
	}

	return base64.StdEncoding.EncodeToString(js), nil
}
