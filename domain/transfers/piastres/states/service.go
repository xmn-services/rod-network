package states

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

// Save saves a state instance
func (app *service) Save(state State) error {
	js, err := app.adapter.ToJSON(state)
	if err != nil {
		return err
	}

	chainHash := state.Chain()
	index := state.Height()
	path := filePath(chainHash, index)
	return app.fileService.Save(path, js)
}
