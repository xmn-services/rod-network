package files

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

type file struct {
	immutable    entities.Immutable
	relativePath string
	chunks       hashtree.HashTree
	amount       uint
}

func createFile(
	immutable entities.Immutable,
	relativePath string,
	chunks hashtree.HashTree,
	amount uint,
) File {
	out := file{
		immutable:    immutable,
		relativePath: relativePath,
		chunks:       chunks,
		amount:       amount,
	}

	return &out
}

// Hash returns the hash
func (obj *file) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// RelativePath returns the relative path
func (obj *file) RelativePath() string {
	return obj.relativePath
}

// Chunks returns the chunks
func (obj *file) Chunks() hashtree.HashTree {
	return obj.chunks
}

// Amount returns the files amount
func (obj *file) Amount() uint {
	return obj.amount
}

// CreatedOn returns the creation time
func (obj *file) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
