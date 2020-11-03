package links

import (
	mined_link "github.com/xmn-services/rod-network/domain/memory/piastres/links/mined"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents the link application
type Application interface {
	Retrieve(hash hash.Hash) (mined_link.Link, error)
}
