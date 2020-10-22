package genesis

import (
	"github.com/xmn-services/rod-network/libs/file"
)

type repository struct {
	adapter         Adapter
	fileRepository  file.Repository
	fileNameWithExt string
}

func createRepository(adapter Adapter, fileRepository file.Repository, fileNameWithExt string) Repository {
	out := repository{
		adapter:         adapter,
		fileRepository:  fileRepository,
		fileNameWithExt: fileNameWithExt,
	}

	return &out
}

// Retrieve retrieves a genesis instance
func (app *repository) Retrieve() (Genesis, error) {
	js, err := app.fileRepository.Retrieve(app.fileNameWithExt)
	if err != nil {
		return nil, err
	}

	return app.adapter.ToGenesis(js)
}
