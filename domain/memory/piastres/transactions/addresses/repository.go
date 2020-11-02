package addresses

import (
	"encoding/json"

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

// Retrieve retrieves an address by hash
func (app *repository) Retrieve(hsh hash.Hash) (Address, error) {
	js, err := app.fileRepository.Retrieve(hsh.String())
	if err != nil {
		return nil, err
	}

	ins := new(JSONAddress)
	err = json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return app.adapter.ToAddress(ins)
}
