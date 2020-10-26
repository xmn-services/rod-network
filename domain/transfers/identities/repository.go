package identities

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

// Retrieve retrieves a identity by hash
func (app *repository) Retrieve(hsh hash.Hash, seed string) (Identity, error) {
	fileName, err := makeFileName(hsh, seed)
	if err != nil {
		return nil, err
	}

	js, err := app.fileRepository.Retrieve(fileName)
	if err != nil {
		return nil, err
	}

	return app.adapter.ToIdentity(js)
}
