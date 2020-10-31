package transactions

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	transfer_transaction "github.com/xmn-services/rod-network/domain/transfers/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/hash"
)

type repository struct {
	builder           Builder
	expenseRepository expenses.Repository
	trRepository      transfer_transaction.Repository
}

func createRepository(
	builder Builder,
	expenseRepository expenses.Repository,
	trRepository transfer_transaction.Repository,
) Repository {
	out := repository{
		builder:           builder,
		expenseRepository: expenseRepository,
		trRepository:      trRepository,
	}

	return &out
}

// Retrieve retrieves a transaction by hash
func (app *repository) Retrieve(hsh hash.Hash) (Transaction, error) {
	trTrx, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	createdOn := trTrx.CreatedOn()
	builder := app.builder.Create().CreatedOn(createdOn)
	if trTrx.HasFees() {
		feesHashes := trTrx.Fees()
		fees, err := app.expenseRepository.RetrieveAll(feesHashes)
		if err != nil {
			return nil, err
		}

		builder.WithFees(fees)
	}

	if trTrx.HasBucket() {
		bucketHash := trTrx.Bucket()
		builder.WithBucket(*bucketHash)
	}

	return builder.Now()
}

// RetrieveAll retrieves all trx from hashes
func (app *repository) RetrieveAll(hashes []hash.Hash) ([]Transaction, error) {
	out := []Transaction{}
	for _, oneHash := range hashes {
		trx, err := app.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		out = append(out, trx)
	}

	return out, nil
}
