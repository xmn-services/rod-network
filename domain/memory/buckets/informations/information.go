package informations

import (
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/buckets/files"
)

type information struct {
	immutable entities.Immutable
	files     []files.File
	mp        map[string]files.File
	parent    Information
}

func createInformation(
	immutable entities.Immutable,
	files []files.File,
	mp map[string]files.File,
) Information {
	return createInformationInternally(immutable, files, mp, nil)
}

func createInformationWithParent(
	immutable entities.Immutable,
	files []files.File,
	mp map[string]files.File,
	parent Information,
) Information {
	return createInformationInternally(immutable, files, mp, parent)
}

func createInformationInternally(
	immutable entities.Immutable,
	files []files.File,
	mp map[string]files.File,
	parent Information,
) Information {
	out := information{
		immutable: immutable,
		files:     files,
		mp:        mp,
		parent:    parent,
	}

	return &out
}

// Hash returns the hash
func (obj *information) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Files returns the files
func (obj *information) Files() []files.File {
	return obj.files
}

// FileByPath returns the file by path
func (obj *information) FileByPath(path string) (files.File, error) {
	if file, ok := obj.mp[path]; ok {
		return file, nil
	}

	str := fmt.Sprintf("the file path (%s) is invalid", path)
	return nil, errors.New(str)
}

// CreatedOn returns the creation time
func (obj *information) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasParent returns true if there is a parent, false otherwise
func (obj *information) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent information, if any
func (obj *information) Parent() Information {
	return obj.parent
}
