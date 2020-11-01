package bucket

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/buckets"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type bucket struct {
	immutable    entities.Immutable
	bucket       buckets.Bucket
	absolutePath string
	pk           encryption.PrivateKey
}

func createBucket(
	immutable entities.Immutable,
	bket buckets.Bucket,
	absolutePath string,
	pk encryption.PrivateKey,
) Bucket {
	out := bucket{
		immutable:    immutable,
		bucket:       bket,
		absolutePath: absolutePath,
		pk:           pk,
	}

	return &out
}

// Hash returns the hash
func (obj *bucket) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Bucket returns the Bucket
func (obj *bucket) Bucket() buckets.Bucket {
	return obj.bucket
}

// AbsolutePath returns the absolutePath
func (obj *bucket) AbsolutePath() string {
	return obj.absolutePath
}

// PrivateKey returns the privateKey
func (obj *bucket) PrivateKey() encryption.PrivateKey {
	return obj.pk
}

// CreatedOn returns the creation time
func (obj *bucket) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
