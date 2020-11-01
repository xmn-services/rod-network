package buckets

import (
	"github.com/xmn-services/rod-network/domain/memory/buckets/files"
	transfer_bucket "github.com/xmn-services/rod-network/domain/transfers/buckets"
	"github.com/xmn-services/rod-network/libs/hash"
)

type repository struct {
	fileRepository files.Repository
	trRepository   transfer_bucket.Repository
	builder        Builder
}

func createRepository(
	fileRepository files.Repository,
	trRepository transfer_bucket.Repository,
	builder Builder,
) Repository {
	out := repository{
		fileRepository: fileRepository,
		trRepository:   trRepository,
		builder:        builder,
	}

	return &out
}

// Retrieve retrieves an bucket instance by hash
func (app *repository) Retrieve(hsh hash.Hash) (Bucket, error) {
	trBucket, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	amount := trBucket.Amount()
	fileHashes := []hash.Hash{}
	leaves := trBucket.Files().Parent().BlockLeaves().Leaves()
	for i := 0; i < int(amount); i++ {
		fileHashes = append(fileHashes, leaves[i].Head())
	}

	files, err := app.fileRepository.RetrieveAll(fileHashes)
	if err != nil {
		return nil, err
	}

	createdOn := trBucket.CreatedOn()
	return app.builder.Create().WithFiles(files).CreatedOn(createdOn).Now()
}
