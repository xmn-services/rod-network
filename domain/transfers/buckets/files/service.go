package files

import (
	libs_file "github.com/xmn-services/rod-network/libs/file"
)

type service struct {
	adapter     Adapter
	fileService libs_file.Service
}

func createService(adapter Adapter, fileService libs_file.Service) Service {
	out := service{
		adapter:     adapter,
		fileService: fileService,
	}

	return &out
}

// Save saves a file instance
func (app *service) Save(file File) error {
	js, err := app.adapter.ToJSON(file)
	if err != nil {
		return err
	}

	return app.fileService.Save(file.Hash().String(), js)
}

// Delete deletes a file instance
func (app *service) Delete(file File) error {
	return app.fileService.Delete(file.Hash().String())
}
