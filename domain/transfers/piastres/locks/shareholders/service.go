package shareholders

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

// Save saves a shareHolder instance
func (app *service) Save(shareHolder ShareHolder) error {
	js, err := app.adapter.ToJSON(shareHolder)
	if err != nil {
		return err
	}

	return app.fileService.Save(shareHolder.Hash().String(), js)
}
