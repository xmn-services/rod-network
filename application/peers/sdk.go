package peers

import "github.com/xmn-services/rod-network/domain/memory/peers"

// Application represents the peer application
type Application interface {
	Retrieve() (peers.Peers, error)
	SaveClear(host string, port uint) error
	SaveOnion(host string, port uint) error
}
