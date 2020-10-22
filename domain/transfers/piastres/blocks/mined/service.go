package mined

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

// Save saves a block instance
func (app *service) Save(block Block) error {
	js, err := app.adapter.ToJSON(block)
	if err != nil {
		return err
	}

	return app.fileService.Save(block.Hash().String(), js)
}
