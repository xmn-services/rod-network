package data

import (
	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

type service struct {
	hashAdapter hash.Adapter
	fileService file.Service
}

func createService(hashAdapter hash.Adapter, fileService file.Service) Service {
	out := service{
		hashAdapter: hashAdapter,
		fileService: fileService,
	}

	return &out
}

// Save saves data instance
func (app *service) Save(data []byte) error {
	hsh, err := app.hashAdapter.FromBytes(data)
	if err != nil {
		return err
	}

	return app.fileService.Save(hsh.String(), data)
}

// Delete deletes data from hash instance
func (app *service) Delete(hsh hash.Hash) error {
	return app.fileService.Delete(hsh.String())
}
