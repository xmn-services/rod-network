package shareholders

import (
	transfer_lock_shareholder "github.com/xmn-services/rod-network/domain/transfers/piastres/locks/shareholders"
)

type service struct {
	adapter    Adapter
	repository Repository
	trService  transfer_lock_shareholder.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	trService transfer_lock_shareholder.Service,
) Service {
	out := service{
		adapter:    adapter,
		repository: repository,
		trService:  trService,
	}

	return &out
}

// Save saves a shareHolder
func (app *service) Save(shareHolder ShareHolder) error {
	hash := shareHolder.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	trHolder, err := app.adapter.ToTransfer(shareHolder)
	if err != nil {
		return err
	}

	return app.trService.Save(trHolder)
}

// SaveAll save a shareHolders
func (app *service) SaveAll(shareHolders []ShareHolder) error {
	for _, oneHolder := range shareHolders {
		err := app.Save(oneHolder)
		if err != nil {
			return err
		}
	}

	return nil
}
