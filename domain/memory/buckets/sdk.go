package buckets

import "github.com/xmn-services/rod-network/domain/memory/buckets/bucket"

// Buckets represents buckets
type Buckets interface {
	All() []bucket.Bucket
}
