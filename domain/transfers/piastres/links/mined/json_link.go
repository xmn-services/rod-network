package mined

import (
	"time"
)

type jsonLink struct {
	Hash      string    `json:"hash"`
	Link      string    `json:"link"`
	Mining    string    `json:"mining"`
	CreatedOn time.Time `json:"created_on"`
}

func createJSONLinkFromLink(ins Link) *jsonLink {
	hash := ins.Hash().String()
	link := ins.Link().String()
	mining := ins.Mining()
	createdOn := ins.CreatedOn()
	return createJSONLink(hash, link, mining, createdOn)
}

func createJSONLink(
	hash string,
	link string,
	mining string,
	createdOn time.Time,
) *jsonLink {
	out := jsonLink{
		Hash:      hash,
		Link:      link,
		Mining:    mining,
		CreatedOn: createdOn,
	}

	return &out
}
