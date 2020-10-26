package statements

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

// Save saves a statement instance
func (app *service) Save(statement Statement) error {
	js, err := app.adapter.ToJSON(statement)
	if err != nil {
		return err
	}

	return app.fileService.Save(statement.Hash().String(), js)
}

// Delete deletes a statement instance
func (app *service) Delete(statement Statement) error {
	return app.fileService.Delete(statement.Hash().String())
}
