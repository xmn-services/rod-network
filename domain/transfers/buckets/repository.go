package buckets

import (
	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

type repository struct {
	adapter        Adapter
	fileRepository file.Repository
}

func createRepository(adapter Adapter, fileRepository file.Repository) Repository {
	out := repository{
		adapter:        adapter,
		fileRepository: fileRepository,
	}

	return &out
}

// RetrieveAll retrieves all buckets
func (app *repository) RetrieveAll() ([]Bucket, error) {
	hashesAsStr, err := app.fileRepository.RetrieveAll(".")
	if err != nil {
		return nil, err
	}

	out := []Bucket{}
	for _, oneHashStr := range hashesAsStr {
		peer, err := app.retrieve(oneHashStr)
		if err != nil {
			return nil, err
		}

		out = append(out, peer)
	}

	return out, nil
}

// Retrieve retrieves a bucket by hash
func (app *repository) Retrieve(hsh hash.Hash) (Bucket, error) {
	return app.retrieve(hsh.String())
}

func (app *repository) retrieve(hashStr string) (Bucket, error) {
	js, err := app.fileRepository.Retrieve(hashStr)
	if err != nil {
		return nil, err
	}

	return app.adapter.ToBucket(js)
}
