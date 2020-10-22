package links

import (
	"time"
)

type jsonLink struct {
	Hash         string    `json:"hash"`
	PreviousLink string    `json:"previous_link"`
	Next         string    `json:"next"`
	Index        uint      `json:"index"`
	CreatedOn    time.Time `json:"created_on"`
}

func createJSONLinkFromLink(ins Link) *jsonLink {
	hash := ins.Hash().String()
	previousLink := ins.PreviousLink().String()
	next := ins.Next().String()
	index := ins.Index()
	createdOn := ins.CreatedOn()
	return createJSONLink(hash, previousLink, next, index, createdOn)
}

func createJSONLink(
	hash string,
	previousLink string,
	next string,
	index uint,
	createdOn time.Time,
) *jsonLink {
	out := jsonLink{
		Hash:         hash,
		PreviousLink: previousLink,
		Next:         next,
		Index:        index,
		CreatedOn:    createdOn,
	}

	return &out
}
