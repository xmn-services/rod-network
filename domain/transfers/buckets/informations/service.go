package informations

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

// Save saves a information instance
func (app *service) Save(information Information) error {
	js, err := app.adapter.ToJSON(information)
	if err != nil {
		return err
	}

	return app.fileService.Save(information.Hash().String(), js)
}

// Delete deletes a information instance
func (app *service) Delete(information Information) error {
	return app.fileService.Delete(information.Hash().String())
}
