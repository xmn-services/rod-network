package buckets

import (
	"github.com/xmn-services/rod-network/domain/memory/buckets"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents the bucket application
type Application interface {
	Add(absolutePath string) error
	Delete(absolutePath string) error
	Retrieve(hash hash.Hash) (buckets.Bucket, error)
}
