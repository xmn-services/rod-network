package chunks

import (
	transfer_chunk "github.com/xmn-services/rod-network/domain/transfers/buckets/files/chunks"
)

type service struct {
	adapter    Adapter
	repository Repository
	trService  transfer_chunk.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	trService transfer_chunk.Service,
) Service {
	out := service{
		adapter:    adapter,
		repository: repository,
		trService:  trService,
	}

	return &out
}

// Save saves a chunk instance
func (app *service) Save(chunk Chunk) error {
	hsh := chunk.Hash()
	_, err := app.repository.Retrieve(hsh)
	if err == nil {
		return nil
	}

	trChunk, err := app.adapter.ToTransfer(chunk)
	if err != nil {
		return err
	}

	return app.trService.Save(trChunk)
}

// SaveAll saves all chunk instances
func (app *service) SaveAll(chunks []Chunk) error {
	for _, oneChunk := range chunks {
		err := app.Save(oneChunk)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete deletes a chunk instance
func (app *service) Delete(chunk Chunk) error {
	trChunk, err := app.adapter.ToTransfer(chunk)
	if err != nil {
		return err
	}

	return app.trService.Delete(trChunk)
}

// DeleteAll deletes all chunk instances
func (app *service) DeleteAll(chunks []Chunk) error {
	for _, oneChunk := range chunks {
		err := app.Delete(oneChunk)
		if err != nil {
			return err
		}
	}

	return nil
}
