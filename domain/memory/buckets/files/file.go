package files

import (
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/buckets/files/chunks"
)

type file struct {
	immutable    entities.Immutable
	relativePath string
	chunks       []chunks.Chunk
	mp           map[string]chunks.Chunk
}

func createFile(
	immutable entities.Immutable,
	relativePath string,
	chunks []chunks.Chunk,
	mp map[string]chunks.Chunk,
) File {
	out := file{
		immutable:    immutable,
		relativePath: relativePath,
		chunks:       chunks,
		mp:           mp,
	}

	return &out
}

// Hash returns the hash
func (obj *file) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// RelativePath returns the relativePath
func (obj *file) RelativePath() string {
	return obj.relativePath
}

// Chunks returns the chunks
func (obj *file) Chunks() []chunks.Chunk {
	return obj.chunks
}

// ChunkByHash returns the chunk by hash
func (obj *file) ChunkByHash(hash hash.Hash) (chunks.Chunk, error) {
	keyname := hash.String()
	if chk, ok := obj.mp[keyname]; ok {
		return chk, nil
	}

	str := fmt.Sprintf("the chunk hash (%s) is invalid", keyname)
	return nil, errors.New(str)
}

// CreatedOn returns the creation time
func (obj *file) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
