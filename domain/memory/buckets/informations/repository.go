package informations

import (
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/buckets/files"
	transfer_information "github.com/xmn-services/rod-network/domain/transfers/buckets/informations"
)

type repository struct {
	fileRepository files.Repository
	trRepository   transfer_information.Repository
	builder        Builder
}

func createRepository(
	fileRepository files.Repository,
	trRepository transfer_information.Repository,
	builder Builder,
) Repository {
	out := repository{
		fileRepository: fileRepository,
		trRepository:   trRepository,
		builder:        builder,
	}

	return &out
}

// Retrieve retrieves an information instance by hash
func (app *repository) Retrieve(hsh hash.Hash) (Information, error) {
	trInformation, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	amount := trInformation.Amount()
	fileHashes := []hash.Hash{}
	leaves := trInformation.Files().Parent().BlockLeaves().Leaves()
	for i := 0; i < int(amount); i++ {
		fileHashes = append(fileHashes, leaves[i].Head())
	}

	files, err := app.fileRepository.RetrieveAll(fileHashes)
	if err != nil {
		return nil, err
	}

	createdOn := trInformation.CreatedOn()
	builder := app.builder.Create().WithFiles(files).CreatedOn(createdOn)
	if trInformation.HasParent() {
		parentHash := trInformation.Parent()
		parent, err := app.Retrieve(*parentHash)
		if err != nil {
			return nil, err
		}

		builder.WithParent(parent)
	}

	return builder.Now()
}
