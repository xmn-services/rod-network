package informations

import (
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/buckets/files"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type information struct {
	immutable entities.Immutable
	files     []files.File
	mp        map[string]files.File
}

func createInformation(
	immutable entities.Immutable,
	files []files.File,
	mp map[string]files.File,
) Information {
	return createInformationInternally(immutable, files, mp)
}

func createInformationInternally(
	immutable entities.Immutable,
	files []files.File,
	mp map[string]files.File,
) Information {
	out := information{
		immutable: immutable,
		files:     files,
		mp:        mp,
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
