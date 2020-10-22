package bucket

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/buckets/bucket/informations"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type bucket struct {
	immutable    entities.Immutable
	information  informations.Information
	absolutePath string
	pk           encryption.PrivateKey
}

func createBucket(
	immutable entities.Immutable,
	information informations.Information,
	absolutePath string,
	pk encryption.PrivateKey,
) Bucket {
	out := bucket{
		immutable:    immutable,
		information:  information,
		absolutePath: absolutePath,
		pk:           pk,
	}

	return &out
}

// Hash returns the hash
func (obj *bucket) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Information returns the information
func (obj *bucket) Information() informations.Information {
	return obj.information
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
