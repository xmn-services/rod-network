package buckets

import (
	application_bucket "github.com/xmn-services/rod-network/application/identities/buckets"
	domain_peers "github.com/xmn-services/rod-network/domain/memory/peers"
)

// Builder represents a bucket client builder
type Builder interface {
	Create() Builder
	WithPeers(peers domain_peers.Peers) Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	Now() (application_bucket.Application, error)
}
