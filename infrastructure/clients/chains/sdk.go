package chains

import (
	application_chain "github.com/xmn-services/rod-network/application/chains"
	domain_peer "github.com/xmn-services/rod-network/domain/memory/peers/peer"
)

// Builder represents a peer client builder
type Builder interface {
	Create() Builder
	WithPeer(peer domain_peer.Peer) Builder
	Now() (application_chain.Application, error)
}
