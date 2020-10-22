package shareholders

import (
	transfer_lock_shareholder "github.com/xmn-services/rod-network/domain/transfers/piastres/locks/shareholders"
)

type adapter struct {
	trBuilder transfer_lock_shareholder.Builder
}

func createAdapter(
	trBuilder transfer_lock_shareholder.Builder,
) Adapter {
	out := adapter{
		trBuilder: trBuilder,
	}

	return &out
}

// ToTransfer converts a shareholder to a transfer shareholder
func (app *adapter) ToTransfer(holder ShareHolder) (transfer_lock_shareholder.ShareHolder, error) {
	hash := holder.Hash()
	key := holder.Key()
	power := holder.Power()
	createdOn := holder.CreatedOn()
	return app.trBuilder.Create().WithHash(hash).WithKey(key).WithPower(power).CreatedOn(createdOn).Now()
}

// ToJSON converts a shareHolder to a jsonShareHolder instance
func (app *adapter) ToJSON(holder ShareHolder) *JSONShareHolder {
	return createJSONShareHolderFromShareHolder(holder)
}

// ToShareHolder converts a json ShareHolder to a shareHolder instance
func (app *adapter) ToShareHolder(ins *JSONShareHolder) (transfer_lock_shareholder.ShareHolder, error) {
	return createShareHolderFromJSON(ins)
}
