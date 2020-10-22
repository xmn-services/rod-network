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

// Save saves a link instance
func (app *service) Save(link Link) error {
	js, err := app.adapter.ToJSON(link)
	if err != nil {
		return err
	}

	return app.fileService.Save(link.Hash().String(), js)
}
