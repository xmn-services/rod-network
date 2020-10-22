package chunks

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

// Save saves a chunk instance
func (app *service) Save(chunk Chunk) error {
	js, err := app.adapter.ToJSON(chunk)
	if err != nil {
		return err
	}

	return app.fileService.Save(chunk.Hash().String(), js)
}

// Delete deletes a chunk instance
func (app *service) Delete(chunk Chunk) error {
	return app.fileService.Delete(chunk.Hash().String())
}
