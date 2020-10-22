package links

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
)

// JSONLink represents a json link
type JSONLink struct {
	PreviousLink string            `json:"previous_link"`
	Next         *blocks.JSONBlock `json:"next"`
	Index        uint              `json:"index"`
	CreatedOn    time.Time         `json:"created_on"`
}

func createJSONLinkFromLink(ins Link) *JSONLink {
	previousLink := ins.PreviousLink().String()
	next := blocks.NewAdapter().ToJSON(ins.Next())
	index := ins.Index()
	createdOn := ins.CreatedOn()
	return createJSONLink(previousLink, next, index, createdOn)
}

func createJSONLink(
	previousLink string,
	next *blocks.JSONBlock,
	index uint,
	createdOn time.Time,
) *JSONLink {
	out := JSONLink{
		PreviousLink: previousLink,
		Next:         next,
		Index:        index,
		CreatedOn:    createdOn,
	}

	return &out
}
