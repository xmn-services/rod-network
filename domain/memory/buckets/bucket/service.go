package bucket

import (
	"github.com/xmn-services/rod-network/domain/memory/buckets/bucket/informations"
	transfer_bucket "github.com/xmn-services/rod-network/domain/transfers/buckets"
)

type service struct {
	adapter            Adapter
	repository         Repository
	informationService informations.Service
	trService          transfer_bucket.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	informationService informations.Service,
	trService transfer_bucket.Service,
) Service {
	out := service{
		adapter:            adapter,
		repository:         repository,
		informationService: informationService,
		trService:          trService,
	}

	return &out
}

// Save saves a bucket instance
func (app *service) Save(bucket Bucket) error {
	path := bucket.AbsolutePath()
	_, err := app.repository.Retrieve(path)
	if err == nil {
		return nil
	}

	information := bucket.Information()
	err = app.informationService.Save(information)
	if err != nil {
		return err
	}

	trBucket, err := app.adapter.ToTransfer(bucket)
	if err != nil {
		return err
	}

	return app.trService.Save(trBucket)

}

// Delete deletes a bucket instance
func (app *service) Delete(bucket Bucket) error {
	information := bucket.Information()
	err := app.informationService.Delete(information)
	if err != nil {
		return err
	}

	trBucket, err := app.adapter.ToTransfer(bucket)
	if err != nil {
		return err
	}

	return app.trService.Delete(trBucket)
}
