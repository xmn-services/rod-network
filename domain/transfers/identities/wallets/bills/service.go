package bills

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

// Save saves a bill instance
func (app *service) Save(bill Bill) error {
	js, err := app.adapter.ToJSON(bill)
	if err != nil {
		return err
	}

	return app.fileService.Save(bill.Hash().String(), js)
}

// Delete deletes a bill instance
func (app *service) Delete(bill Bill) error {
	return app.fileService.Delete(bill.Hash().String())
}
