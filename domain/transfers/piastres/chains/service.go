package chains

import (
	"github.com/xmn-services/rod-network/libs/file"
)

type service struct {
	adapter         Adapter
	fileService     file.Service
	fileNameWithExt string
}

func createService(adapter Adapter, fileService file.Service, fileNameWithExt string) Service {
	out := service{
		adapter:         adapter,
		fileService:     fileService,
		fileNameWithExt: fileNameWithExt,
	}

	return &out
}

// Save saves a chain instance
func (app *service) Save(chain Chain) error {
	js, err := app.adapter.ToJSON(chain)
	if err != nil {
		return err
	}

	return app.fileService.Save(app.fileNameWithExt, js)
}
