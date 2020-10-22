package blocks

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	transfer_block "github.com/xmn-services/rod-network/domain/transfers/piastres/blocks"
)

type service struct {
	adapter    Adapter
	repository Repository
	trxService transactions.Service
	trService  transfer_block.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	trxService transactions.Service,
	trService transfer_block.Service,
) Service {
	out := service{
		adapter:    adapter,
		repository: repository,
		trxService: trxService,
		trService:  trService,
	}

	return &out
}

// Save saves a block
func (app *service) Save(block Block) error {
	hash := block.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	trx := block.Transactions()
	err = app.trxService.SaveAll(trx)
	if err != nil {
		return err
	}

	trBlock, err := app.adapter.ToTransfer(block)
	if err != nil {
		return err
	}

	return app.trService.Save(trBlock)
}
