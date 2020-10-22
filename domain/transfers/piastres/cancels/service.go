package cancels

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

// Save saves a cancel instance
func (app *service) Save(cancel Cancel) error {
	js, err := app.adapter.ToJSON(cancel)
	if err != nil {
		return err
	}

	return app.fileService.Save(cancel.Hash().String(), js)
}
