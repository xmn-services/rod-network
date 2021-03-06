package locks

import (
	"github.com/xmn-services/rod-network/libs/file"
)

type service struct {
	adapter     Adapter
	fileService file.Service
}

func createService(adapter Adapter, fileService file.Service) Service {
	out := service{
		adapter:     adapter,
		fileService: fileService,
	}

	return &out
}

// Save saves a lock instance
func (app *service) Save(lock Lock) error {
	js, err := app.adapter.ToJSON(lock)
	if err != nil {
		return err
	}

	return app.fileService.Save(lock.Hash().String(), js)
}
