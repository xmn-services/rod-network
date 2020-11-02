package addresses

import (
	"encoding/json"

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

// Save saves an address instance
func (app *service) Save(address Address) error {
	jsIns := app.adapter.ToJSON(address)
	js, err := json.Marshal(jsIns)
	if err != nil {
		return err
	}

	return app.fileService.Save(address.Hash().String(), js)
}
