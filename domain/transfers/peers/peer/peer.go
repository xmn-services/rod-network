package peer

type peer struct {
	host    string
	port    uint
	isClear bool
	isOnion bool
}

func createPeerWithClear(
	host string,
	port uint,
) Peer {
	return createPeerInternally(host, port, true, false)
}

func createPeerWithOnion(
	host string,
	port uint,
) Peer {
	return createPeerInternally(host, port, false, true)
}

func createPeerInternally(
	host string,
	port uint,
	isClear bool,
	isOnion bool,
) Peer {
	out := peer{
		host:    host,
		port:    port,
		isClear: isClear,
		isOnion: isOnion,
	}

	return &out
}

// Host returns the host
func (obj *peer) Host() string {
	return obj.host
}

// Port returns the port
func (obj *peer) Port() uint {
	return obj.port
}

// IsClear returns true if the peer is clear, false otherwise
func (obj *peer) IsClear() bool {
	return obj.isClear
}

// IsOnion returns true if the peer is onion, false otherwise
func (obj *peer) IsOnion() bool {
	return obj.isOnion
}
