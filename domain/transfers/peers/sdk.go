package peers

import (
	"github.com/xmn-services/rod-network/domain/transfers/peers/peer"
	"github.com/xmn-services/rod-network/libs/file"
)

// NewService creates a new service instance
func NewService(fileService file.Service, filePathWithName string) Service {
	adapter := NewAdapter()
	return createService(adapter, fileService, filePathWithName)
}

// NewRepository creates a new repository instance
func NewRepository(fileRepository file.Repository, filePathWithName string) Repository {
	adapter := NewAdapter()
	return createRepository(adapter, fileRepository, filePathWithName)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	peerAdapter := peer.NewAdapter()
	builder := NewBuilder()
	return createAdapter(peerAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a peers adapter
type Adapter interface {
	ToPeers(urls []string) (Peers, error)
	ToURLs(peers Peers) []string
}

// Builder represents peers builder
type Builder interface {
	Create() Builder
	WithPeers(peers []peer.Peer) Builder
	Now() (Peers, error)
}

// Peers represents peers
type Peers interface {
	All() []peer.Peer
}

// Repository represents a peers repository
type Repository interface {
	Retrieve() (Peers, error)
}

// Service represents a peer service
type Service interface {
	Save(peers Peers) error
}
