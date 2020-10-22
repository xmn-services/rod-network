package chunks

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type chunk struct {
	immutable   entities.Immutable
	sizeInBytes uint
	data        hash.Hash
}

func createChunk(
	immutable entities.Immutable,
	sizeInBytes uint,
	data hash.Hash,
) Chunk {
	out := chunk{
		immutable:   immutable,
		sizeInBytes: sizeInBytes,
		data:        data,
	}

	return &out
}

// Hash returns the hash
func (obj *chunk) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// SizeInBytes returns the sizeInBytes
func (obj *chunk) SizeInBytes() uint {
	return obj.sizeInBytes
}

// Data returns the data hash
func (obj *chunk) Data() hash.Hash {
	return obj.data
}

// CreatedOn returns the creation time
func (obj *chunk) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
