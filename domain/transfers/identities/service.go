package identities

import (
	"github.com/xmn-services/rod-network/libs/file"
)

type service struct {
	adapter     Adapter
	repository  Repository
	fileService file.Service
}

func createService(adapter Adapter, repository Repository, fileService file.Service) Service {
	out := service{
		adapter:     adapter,
		repository:  repository,
		fileService: fileService,
	}

	return &out
}

// Save saves a identity instance
func (app *service) Save(identity Identity) error {
	fileName, err := makeFileName(identity.Hash(), identity.Seed())
	if err != nil {
		return err
	}

	js, err := app.adapter.ToJSON(identity)
	if err != nil {
		return err
	}

	return app.fileService.Save(fileName, js)
}

// Delete deletes a identity instance
func (app *service) Delete(identity Identity) error {
	fileName, err := makeFileName(identity.Hash(), identity.Seed())
	if err != nil {
		return err
	}

	return app.fileService.Delete(fileName)
}
