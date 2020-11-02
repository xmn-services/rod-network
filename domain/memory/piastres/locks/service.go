package locks

import (
	transfer_lock "github.com/xmn-services/rod-network/domain/transfers/piastres/locks"
)

type service struct {
	adapter    Adapter
	repository Repository
	trService  transfer_lock.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	trService transfer_lock.Service,
) Service {
	out := service{
		adapter:    adapter,
		repository: repository,
		trService:  trService,
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

	trLock, err := app.adapter.ToTransfer(lock)
	if err != nil {
		return err
	}

	return app.trService.Save(trLock)
}
