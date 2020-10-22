package chunks

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

// Retrieve retrieves a chunk by hash
func (app *repository) Retrieve(hsh hash.Hash) (Chunk, error) {
	js, err := app.fileRepository.Retrieve(hsh.String())
	if err != nil {
		return nil, err
	}

	return app.adapter.ToChunk(js)
}

// RetrieveAll retrieves chunks by hashes
func (app *repository) RetrieveAll(hashes []hash.Hash) ([]Chunk, error) {
	out := []Chunk{}
	for _, oneHash := range hashes {
		chunk, err := app.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		out = append(out, chunk)
	}

	return out, nil
}
