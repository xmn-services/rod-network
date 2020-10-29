package reserves

import (
	"github.com/xmn-services/rod-network/domain/memory/exchanges"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Reserve represents a request reserve
type Reserve interface {
	entities.Immutable
	Content() Content
	Signature() signature.RingSignature
}

// Content represents a reserve content
type Content interface {
	entities.Immutable
	Request() requests.Request
	Nonces() []uint
}
