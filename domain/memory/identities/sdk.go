package identities

import (
	"github.com/xmn-services/rod-network/domain/memory/buckets"
	"github.com/xmn-services/rod-network/domain/memory/identities/wallets"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Identity represents the identity
type Identity interface {
	entities.Mutable
	Seed() string
	Name() string
	Root() string
	Wallet() wallets.Wallet
	HasBuckets() bool
	Buckets() buckets.Bucket
}
