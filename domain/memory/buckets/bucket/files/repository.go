package files

import (
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/buckets/bucket/files/chunks"
	transfer_file "github.com/xmn-services/rod-network/domain/transfers/buckets/files"
)

type repository struct {
	chunkRepository chunks.Repository
	trRepository    transfer_file.Repository
	builder         Builder
}

func createRepository(
	chunkRepository chunks.Repository,
	trRepository transfer_file.Repository,
	builder Builder,
) Repository {
	out := repository{
		chunkRepository: chunkRepository,
		trRepository:    trRepository,
		builder:         builder,
	}

	return &out
}

// Retrieve retrieves a file instance by hash
func (app *repository) Retrieve(hsh hash.Hash) (File, error) {
	trFile, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	amount := trFile.Amount()
	chunkHashes := []hash.Hash{}
	leaves := trFile.Chunks().Parent().BlockLeaves().Leaves()
	for i := 0; i < int(amount); i++ {
		chunkHashes = append(chunkHashes, leaves[i].Head())
	}

	chunks, err := app.chunkRepository.RetrieveAll(chunkHashes)
	if err != nil {
		return nil, err
	}

	sizeInBytes := trFile.RelativePath()
	createdOn := trFile.CreatedOn()
	return app.builder.Create().WithRelativePath(sizeInBytes).WithChunks(chunks).CreatedOn(createdOn).Now()
}

// RetrieveAll retrieves all file instances by hashes
func (app *repository) RetrieveAll(hashes []hash.Hash) ([]File, error) {
	out := []File{}
	for _, oneHash := range hashes {
		chk, err := app.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		out = append(out, chk)
	}

	return out, nil
}
