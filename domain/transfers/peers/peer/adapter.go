package peer

import (
	"fmt"
	"net/url"
	"strconv"
)

type adapter struct {
	builder Builder
}

func createAdapter(
	builder Builder,
) Adapter {
	out := adapter{
		builder: builder,
	}

	return &out
}

// ToPeer converts a rawURL to peer
func (app *adapter) ToPeer(rawURL string) (Peer, error) {
	url, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	host := url.Hostname()
	portStr := url.Port()
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithHost(host).WithPort(uint(port))
	if url.Scheme == clearProtocol {
		builder.IsClear()
	}

	if url.Scheme == onionProtocol {
		builder.IsOnion()
	}

	return builder.Now()
}

// ToURL converts a peer to url
func (app *adapter) ToURL(peer Peer) string {
	protocol := ""
	if peer.IsClear() {
		protocol = clearProtocol
	}

	if peer.IsOnion() {
		protocol = onionProtocol
	}

	host := peer.Host()
	port := peer.Port()
	return fmt.Sprintf("%s://%s:%d", protocol, host, port)
}
