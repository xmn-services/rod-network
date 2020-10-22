package entities

import (
	"time"

	"github.com/xmn-services/rod-network/libs/hash"
)

type immutable struct {
	hash      hash.Hash
	createdOn time.Time
}

func createImmutable(hash hash.Hash, createdOn time.Time) Immutable {
	out := immutable{
		hash:      hash,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *immutable) Hash() hash.Hash {
	return obj.hash
}

// CreatedOn returns the creation time
func (obj *immutable) CreatedOn() time.Time {
	return obj.createdOn
}
