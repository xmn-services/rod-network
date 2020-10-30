package transactions

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/cancels"
	"github.com/xmn-services/rod-network/libs/hash"
)

type element struct {
	cancel cancels.Cancel
	bucket *hash.Hash
}

func createElementWithCancel(
	cancel cancels.Cancel,
) Element {
	return createElementInternally(cancel, nil)
}

func createElementWithBucket(
	bucket *hash.Hash,
) Element {
	return createElementInternally(nil, bucket)
}

func createElementInternally(
	cancel cancels.Cancel,
	bucket *hash.Hash,
) Element {
	out := element{
		cancel: cancel,
		bucket: bucket,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	if obj.IsBucket() {
		return *obj.bucket
	}

	return obj.Cancel().Hash()
}

// IsCancel returns true if there is a cancel, false otherwise
func (obj *element) IsCancel() bool {
	return obj.cancel != nil
}

// Cancel returns the cancel, if any
func (obj *element) Cancel() cancels.Cancel {
	return obj.cancel
}

// IsBucket returns true if there is a bucket, false otherwise
func (obj *element) IsBucket() bool {
	return obj.bucket != nil
}

// Bucket returns the bucket, if any
func (obj *element) Bucket() *hash.Hash {
	return obj.bucket
}
