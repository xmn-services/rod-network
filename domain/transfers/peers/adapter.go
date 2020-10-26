package peers

import "github.com/xmn-services/rod-network/domain/transfers/peers/peer"

type adapter struct {
	peerAdapter peer.Adapter
	builder     Builder
}

func createAdapter(
	peerAdapter peer.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		peerAdapter: peerAdapter,
		builder:     builder,
	}

	return &out
}

// ToPeers converts urls to peers
func (app *adapter) ToPeers(urls []string) (Peers, error) {
	return nil, nil
}

// ToURLs converts peers to urls
func (app *adapter) ToURLs(peers Peers) []string {
	urls := []string{}
	all := peers.All()
	for _, onePeer := range all {
		url := app.peerAdapter.ToURL(onePeer)
		urls = append(urls, url)
	}

	return urls
}
