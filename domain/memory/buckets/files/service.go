package files

import (
	"github.com/xmn-services/rod-network/domain/memory/buckets/files/chunks"
	transfer_file "github.com/xmn-services/rod-network/domain/transfers/buckets/files"
)

type service struct {
	adapter      Adapter
	repository   Repository
	chunkService chunks.Service
	trService    transfer_file.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	chunkService chunks.Service,
	trService transfer_file.Service,
) Service {
	out := service{
		adapter:      adapter,
		repository:   repository,
		chunkService: chunkService,
		trService:    trService,
	}

	return &out
}

// Save saves a file instance
func (app *service) Save(file File) error {
	hash := file.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	chunks := file.Chunks()
	err = app.chunkService.SaveAll(chunks)
	if err != nil {
		return err
	}

	trFile, err := app.adapter.ToTransfer(file)
	if err != nil {
		return err
	}

	return app.trService.Save(trFile)
}

// SaveAll saves all chunk instances
func (app *service) SaveAll(files []File) error {
	for _, oneFile := range files {
		err := app.Save(oneFile)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete deletes a chunk instance
func (app *service) Delete(file File) error {
	trChunk, err := app.adapter.ToTransfer(file)
	if err != nil {
		return err
	}

	return app.trService.Delete(trChunk)
}

// DeleteAll deletes all chunk instances
func (app *service) DeleteAll(files []File) error {
	for _, oneFile := range files {
		err := app.Delete(oneFile)
		if err != nil {
			return err
		}
	}

	return nil
}
