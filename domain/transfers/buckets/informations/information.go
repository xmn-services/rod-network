package informations

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

type information struct {
	immutable entities.Immutable
	files     hashtree.HashTree
	amount    uint
}

func createInformation(
	immutable entities.Immutable,
	files hashtree.HashTree,
	amount uint,
) Information {
	return createInformationInternally(immutable, files, amount)
}

func createInformationInternally(
	immutable entities.Immutable,
	files hashtree.HashTree,
	amount uint,
) Information {
	out := information{
		immutable: immutable,
		files:     files,
		amount:    amount,
	}

	return &out
}

// Hash returns the hash
func (obj *information) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Files return the files
func (obj *information) Files() hashtree.HashTree {
	return obj.files
}

// Amount returns the amount
func (obj *information) Amount() uint {
	return obj.amount
}

// CreatedOn returns the creation time
func (obj *information) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
