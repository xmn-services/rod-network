package peer

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
	String() string
}
