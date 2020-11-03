package buckets

import (
	public_bucket "github.com/xmn-services/rod-network/domain/memory/buckets"
	"github.com/xmn-services/rod-network/domain/memory/identities/buckets/bucket"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Builder represents the bucket builder
type Builder interface {
	Create() Builder
	WithBuckets(buckets []bucket.Bucket) Builder
	WithFollows(follows []hash.Hash) Builder
	Now() (Buckets, error)
}

// Buckets represents buckets
type Buckets interface {
	All() []bucket.Bucket
	HasFollows() bool
	Follows() []hash.Hash
	Fetch(absoluteFilePath string) (bucket.Bucket, error)
	Add(bucket bucket.Bucket) error
	Delete(bucket bucket.Bucket) error
	Follow(bucket public_bucket.Bucket) error
}
