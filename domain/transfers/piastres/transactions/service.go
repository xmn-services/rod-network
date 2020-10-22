package transactions

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

// Save saves a transaction instance
func (app *service) Save(transaction Transaction) error {
	js, err := app.adapter.ToJSON(transaction)
	if err != nil {
		return err
	}

	return app.fileService.Save(transaction.Hash().String(), js)
}
