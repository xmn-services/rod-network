package buckets

import "github.com/xmn-services/rod-network/domain/memory/identities/buckets/bucket"

// Builder represents the bucket builder
type Builder interface {
	Create() Builder
	WithBuckets(buckets []bucket.Bucket) Builder
	Now() (Buckets, error)
}

// Buckets represents buckets
type Buckets interface {
	All() []bucket.Bucket
	Fetch(absoluteFilePath string) (bucket.Bucket, error)
	Add(bucket bucket.Bucket) error
	Delete(bucket bucket.Bucket) error
}
