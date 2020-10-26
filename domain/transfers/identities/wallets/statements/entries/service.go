package entries

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

// Save saves a entry instance
func (app *service) Save(entry Entry) error {
	js, err := app.adapter.ToJSON(entry)
	if err != nil {
		return err
	}

	return app.fileService.Save(entry.Hash().String(), js)
}

// Delete deletes a entry instance
func (app *service) Delete(entry Entry) error {
	return app.fileService.Delete(entry.Hash().String())
}
