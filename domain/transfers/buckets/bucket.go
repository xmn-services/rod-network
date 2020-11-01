package buckets

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

type bucket struct {
	immutable entities.Immutable
	files     hashtree.HashTree
	amount    uint
}

func createBucket(
	immutable entities.Immutable,
	files hashtree.HashTree,
	amount uint,
) Bucket {
	return createBucketInternally(immutable, files, amount)
}

func createBucketInternally(
	immutable entities.Immutable,
	files hashtree.HashTree,
	amount uint,
) Bucket {
	out := bucket{
		immutable: immutable,
		files:     files,
		amount:    amount,
	}

	return &out
}

// Hash returns the hash
func (obj *bucket) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Files return the files
func (obj *bucket) Files() hashtree.HashTree {
	return obj.files
}

// Amount returns the amount
func (obj *bucket) Amount() uint {
	return obj.amount
}

// CreatedOn returns the creation time
func (obj *bucket) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
