package buckets

import (
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/buckets/files"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type bucket struct {
	immutable entities.Immutable
	files     []files.File
	mp        map[string]files.File
}

func createBucket(
	immutable entities.Immutable,
	files []files.File,
	mp map[string]files.File,
) Bucket {
	return createBucketInternally(immutable, files, mp)
}

func createBucketInternally(
	immutable entities.Immutable,
	files []files.File,
	mp map[string]files.File,
) Bucket {
	out := bucket{
		immutable: immutable,
		files:     files,
		mp:        mp,
	}

	return &out
}

// Hash returns the hash
func (obj *bucket) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Files returns the files
func (obj *bucket) Files() []files.File {
	return obj.files
}

// FileByPath returns the file by path
func (obj *bucket) FileByPath(path string) (files.File, error) {
	if file, ok := obj.mp[path]; ok {
		return file, nil
	}

	str := fmt.Sprintf("the file path (%s) is invalid", path)
	return nil, errors.New(str)
}

// CreatedOn returns the creation time
func (obj *bucket) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
