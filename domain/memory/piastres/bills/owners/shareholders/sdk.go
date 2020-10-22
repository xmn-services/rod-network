package shareholders

import (
	"hash"

	"github.com/xmn-services/rod-network/libs/entities"
)

// ShareHolder represents an owner's shareholder
type ShareHolder interface {
	entities.Immutable
	Pointer() hash.Hash
	Name() string
	Description() string
}
