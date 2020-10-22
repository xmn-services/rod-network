package data

import (
	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

type repository struct {
	fileRepository file.Repository
}

func createRepository(fileRepository file.Repository) Repository {
	out := repository{
		fileRepository: fileRepository,
	}

	return &out
}

// Retrieve retrieves data by hash
func (app *repository) Retrieve(hsh hash.Hash) ([]byte, error) {
	return app.fileRepository.Retrieve(hsh.String())
}
