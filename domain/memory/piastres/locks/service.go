package locks

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
	transfer_lock "github.com/xmn-services/rod-network/domain/transfers/piastres/locks"
)

type service struct {
	adapter            Adapter
	repository         Repository
	shareHolderService shareholders.Service
	trService          transfer_lock.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	shareHolderService shareholders.Service,
	trService transfer_lock.Service,
) Service {
	out := service{
		adapter:            adapter,
		repository:         repository,
		shareHolderService: shareHolderService,
		trService:          trService,
	}

	return &out
}

// Save saves a lock instance
func (app *service) Save(lock Lock) error {
	hash := lock.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	holders := lock.ShareHolders()
	err = app.shareHolderService.SaveAll(holders)
	if err != nil {
		return err
	}

	trLock, err := app.adapter.ToTransfer(lock)
	if err != nil {
		return err
	}

	return app.trService.Save(trLock)
}
