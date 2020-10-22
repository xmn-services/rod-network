package buckets

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

// Save saves a bucket instance
func (app *service) Save(bucket Bucket) error {
	js, err := app.adapter.ToJSON(bucket)
	if err != nil {
		return err
	}

	return app.fileService.Save(bucket.Hash().String(), js)
}

// Delete deletes a bucket instance
func (app *service) Delete(bucket Bucket) error {
	return app.fileService.Delete(bucket.Hash().String())
}
