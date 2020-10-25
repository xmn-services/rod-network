package peers

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/peers/peer"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type peers struct {
	mutable entities.Mutable
	peers   []peer.Peer
}

func createPeers(
	mutable entities.Mutable,
	lst []peer.Peer,
) Peers {
	out := peers{
		mutable: mutable,
		peers:   lst,
	}

	return &out
}

// Hash returns the hash
func (obj *peers) Hash() hash.Hash {
	return obj.mutable.Hash()
}

// All returns the peers
func (obj *peers) All() []peer.Peer {
	return obj.peers
}

// LastUpdatedOn returns the lastUpdatedOn time
func (obj *peers) LastUpdatedOn() time.Time {
	return obj.mutable.LastUpdatedOn()
}

// CreatedOn returns the creation time
func (obj *peers) CreatedOn() time.Time {
	return obj.mutable.CreatedOn()
}
