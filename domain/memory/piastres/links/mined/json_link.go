package mined

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/links"
)

// JSONLink represents a json link
type JSONLink struct {
	Link      *links.JSONLink `json:"link"`
	Mining    string          `json:"mining"`
	CreatedOn time.Time       `json:"created_on"`
}

func createJSONLinkFromLink(ins Link) *JSONLink {
	link := links.NewAdapter().ToJSON(ins.Link())
	mining := ins.Mining()
	createdOn := ins.CreatedOn()
	return createJSONLink(link, mining, createdOn)
}

func createJSONLink(
	link *links.JSONLink,
	mining string,
	createdOn time.Time,
) *JSONLink {
	out := JSONLink{
		Link:      link,
		Mining:    mining,
		CreatedOn: createdOn,
	}

	return &out
}
