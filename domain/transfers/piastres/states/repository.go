package states

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

// Retrieve retrieves a state by chain and index
func (app *repository) Retrieve(chainHash hash.Hash, height uint) (State, error) {
	path := filePath(chainHash, height)
	js, err := app.fileRepository.Retrieve(path)
	if err != nil {
		return nil, err
	}

	return app.adapter.ToState(js)
}
