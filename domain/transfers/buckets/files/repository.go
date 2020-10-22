package files

import (
	libs_file "github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

type repository struct {
	adapter        Adapter
	fileRepository libs_file.Repository
}

func createRepository(adapter Adapter, fileRepository libs_file.Repository) Repository {
	out := repository{
		adapter:        adapter,
		fileRepository: fileRepository,
	}

	return &out
}

// Retrieve retrieves a file by hash
func (app *repository) Retrieve(hsh hash.Hash) (File, error) {
	js, err := app.fileRepository.Retrieve(hsh.String())
	if err != nil {
		return nil, err
	}

	return app.adapter.ToFile(js)
}

// RetrieveAll retrieves files by hashes
func (app *repository) RetrieveAll(hashes []hash.Hash) ([]File, error) {
	out := []File{}
	for _, oneHash := range hashes {
		file, err := app.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		out = append(out, file)
	}

	return out, nil
}
