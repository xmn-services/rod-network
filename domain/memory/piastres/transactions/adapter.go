package transactions

import (
	transfer_transaction "github.com/xmn-services/rod-network/domain/transfers/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/hash"
)

type adapter struct {
	trBuilder transfer_transaction.Builder
}

func createAdapter(
	trBuilder transfer_transaction.Builder,
) Adapter {
	out := adapter{
		trBuilder: trBuilder,
	}

	return &out
}

// ToTransfer converts a transaction to a transfer transaction instance
func (app *adapter) ToTransfer(trx Transaction) (transfer_transaction.Transaction, error) {
	hsh := trx.Hash()
	content := trx.Content()
	triggersOn := content.TriggersOn()
	sig := trx.Signature()
	createdOn := trx.CreatedOn()

	builder := app.trBuilder.Create().
		WithHash(hsh).
		TriggersOn(triggersOn).
		WithSignature(sig).
		CreatedOn(createdOn)

	if content.HasFees() {
		fees := content.Fees()
		feeHashes := []hash.Hash{}
		for _, oneFee := range fees {
			feeHashes = append(feeHashes, oneFee.Hash())
		}

		builder.WithFees(feeHashes)
	}

	if content.HasElement() {
		element := content.Element()
		if element.IsBucket() {
			bucket := element.Bucket()
			builder.WithBucket(*bucket)
		}

		if element.IsCancel() {
			cancel := element.Cancel().Hash()
			builder.WithCancel(cancel)
		}
	}

	return builder.Now()
}

// ToJSON converts a transaction to a JSON instances
func (app *adapter) ToJSON(trx Transaction) *JSONTransaction {
	return createJSONTransactionFromTransaction(trx)
}

// ToTransaction converts a JSON transaction to a Transaction instances
func (app *adapter) ToTransaction(ins *JSONTransaction) (Transaction, error) {
	return createTransactionFromJSON(ins)
}
