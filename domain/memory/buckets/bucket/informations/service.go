package informations

import (
	"github.com/xmn-services/rod-network/domain/memory/buckets/bucket/files"
	transfer_information "github.com/xmn-services/rod-network/domain/transfers/buckets/informations"
)

type service struct {
	adapter     Adapter
	repository  Repository
	fileService files.Service
	trService   transfer_information.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	fileService files.Service,
	trService transfer_information.Service,
) Service {
	out := service{
		adapter:     adapter,
		repository:  repository,
		fileService: fileService,
		trService:   trService,
	}

	return &out
}

// Save saves an information instance
func (app *service) Save(information Information) error {
	hash := information.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	files := information.Files()
	err = app.fileService.SaveAll(files)
	if err != nil {
		return err
	}

	if information.HasParent() {
		parent := information.Parent()
		err := app.Save(parent)
		if err != nil {
			return err
		}
	}

	trInformation, err := app.adapter.ToTransfer(information)
	if err != nil {
		return err
	}

	return app.trService.Save(trInformation)
}

// Delete deletes an information instance
func (app *service) Delete(information Information) error {
	files := information.Files()
	err := app.fileService.DeleteAll(files)
	if err != nil {
		return err
	}

	if information.HasParent() {
		parent := information.Parent()
		err := app.Delete(parent)
		if err != nil {
			return err
		}
	}

	trInformation, err := app.adapter.ToTransfer(information)
	if err != nil {
		return err
	}

	return app.trService.Delete(trInformation)
}
