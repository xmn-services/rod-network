package peer

const clearProtocol = "https"
const onionProtocol = "onion"

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	return createAdapter(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a peer adapter
type Adapter interface {
	ToPeer(rawURL string) (Peer, error)
	ToURL(peer Peer) string
}

// Builder represents a peer builder
type Builder interface {
	Create() Builder
	WithHost(host string) Builder
	WithPort(port uint) Builder
	IsClear() Builder
	IsOnion() Builder
	Now() (Peer, error)
}

// Peer represents a peer
type Peer interface {
	Host() string
	Port() uint
	IsClear() bool
	IsOnion() bool
}
