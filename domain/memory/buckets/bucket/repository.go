package bucket

import (
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/buckets/bucket/informations"
	transfer_bucket "github.com/xmn-services/rod-network/domain/transfers/buckets"
)

type repository struct {
	hashAdapter           hash.Adapter
	informationRepository informations.Repository
	trRepository          transfer_bucket.Repository
	builder               Builder
}

func createRepository(
	hashAdapter hash.Adapter,
	informationRepository informations.Repository,
	trRepository transfer_bucket.Repository,
	builder Builder,
) Repository {
	out := repository{
		hashAdapter:           hashAdapter,
		informationRepository: informationRepository,
		trRepository:          trRepository,
		builder:               builder,
	}

	return &out
}

// RetrieveAll retrieves all buckets
func (app *repository) RetrieveAll() ([]Bucket, error) {
	trBuckets, err := app.trRepository.RetrieveAll()
	if err != nil {
		return nil, err
	}

	out := []Bucket{}
	for _, oneTrBucket := range trBuckets {
		bucket, err := app.build(oneTrBucket)
		if err != nil {
			return nil, err
		}

		out = append(out, bucket)
	}

	return out, nil
}

// Retrieve retrieves a bucket by AbsolutePath
func (app *repository) Retrieve(absolutePath string) (Bucket, error) {
	hsh, err := app.hashAdapter.FromBytes([]byte(absolutePath))
	if err != nil {
		return nil, err
	}

	trBucket, err := app.trRepository.Retrieve(*hsh)
	if err != nil {
		return nil, err
	}

	return app.build(trBucket)
}

func (app *repository) build(trBucket transfer_bucket.Bucket) (Bucket, error) {
	informationHash := trBucket.Information()
	information, err := app.informationRepository.Retrieve(informationHash)
	if err != nil {
		return nil, err
	}

	path := trBucket.AbsolutePath()
	pk := trBucket.PrivateKey()
	createdOn := trBucket.CreatedOn()
	return app.builder.Create().
		WithAbsolutePath(path).
		WithInformation(information).
		WithPrivateKey(pk).
		CreatedOn(createdOn).Now()
}
