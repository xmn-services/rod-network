package locks

import (
	transfer_lock "github.com/xmn-services/rod-network/domain/transfers/piastres/locks"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

type adapter struct {
	htBuilder hashtree.Builder
	trBuilder transfer_lock.Builder
}

func createAdapter(
	htBuilder hashtree.Builder,
	trBuilder transfer_lock.Builder,
) Adapter {
	out := adapter{
		htBuilder: htBuilder,
		trBuilder: trBuilder,
	}

	return &out
}

// ToTransfer converts a lock to a transfer lock
func (app *adapter) ToTransfer(lock Lock) (transfer_lock.Lock, error) {
	hsh := lock.Hash()
	pubKeys := lock.PublicKeys()
	createdOn := lock.CreatedOn()
	return app.trBuilder.Create().
		WithHash(hsh).
		WithPublicKeys(pubKeys).
		CreatedOn(createdOn).
		Now()
}

// ToJSON converts a lock to a json Lock instance
func (app *adapter) ToJSON(ins Lock) *JSONLock {
	return createJSONLockFromLock(ins)
}

// ToLock converts a lock instance from json lock
func (app *adapter) ToLock(ins *JSONLock) (Lock, error) {
	return createLockFromJSON(ins)
}
