package buckets

import (
	"io/ioutil"
	"math"
	"path/filepath"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/buckets"
	"github.com/xmn-services/rod-network/domain/memory/buckets/files"
	"github.com/xmn-services/rod-network/domain/memory/buckets/files/chunks"
	"github.com/xmn-services/rod-network/domain/memory/identities"
	identity_buckets "github.com/xmn-services/rod-network/domain/memory/identities/buckets/bucket"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/hash"
)

type application struct {
	hashAdapter           hash.Adapter
	pkFactory             encryption.Factory
	chunkBuilder          chunks.Builder
	fileBuilder           files.Builder
	bucketBuilder         buckets.Builder
	bucketRepository      buckets.Repository
	bucketService         buckets.Service
	identityBucketBuilder identity_buckets.Builder
	identityRepository    identities.Repository
	identityService       identities.Service
	identityBuilder       identities.Builder
	name                  string
	password              string
	seed                  string
	chunkSizeInBytes      uint
}

func createApplication(
	hashAdapter hash.Adapter,
	pkFactory encryption.Factory,
	chunkBuilder chunks.Builder,
	fileBuilder files.Builder,
	bucketBuilder buckets.Builder,
	bucketRepository buckets.Repository,
	bucketService buckets.Service,
	identityBucketBuilder identity_buckets.Builder,
	identityRepository identities.Repository,
	identityService identities.Service,
	identityBuilder identities.Builder,
	name string,
	password string,
	seed string,
	chunkSizeInBytes uint,
) Application {
	out := application{
		hashAdapter:           hashAdapter,
		pkFactory:             pkFactory,
		chunkBuilder:          chunkBuilder,
		fileBuilder:           fileBuilder,
		bucketBuilder:         bucketBuilder,
		bucketRepository:      bucketRepository,
		bucketService:         bucketService,
		identityBucketBuilder: identityBucketBuilder,
		identityRepository:    identityRepository,
		identityService:       identityService,
		identityBuilder:       identityBuilder,
		name:                  name,
		password:              password,
		seed:                  seed,
		chunkSizeInBytes:      chunkSizeInBytes,
	}

	return &out
}

// Add adds the bucket path
func (app *application) Add(absolutePath string) error {
	files, err := app.dirToFiles(absolutePath, ".")
	if err != nil {
		return err
	}

	bucketCreatedOn := time.Now().UTC()
	bucket, err := app.bucketBuilder.Create().WithFiles(files).CreatedOn(bucketCreatedOn).Now()
	if err != nil {
		return err
	}

	pk, err := app.pkFactory.Create()
	if err != nil {
		return err
	}

	createdOn := time.Now().UTC()
	identityBucket, err := app.identityBucketBuilder.Create().
		WithBucket(bucket).
		WithAbsolutePath(absolutePath).
		WithPrivateKey(pk).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		return err
	}

	// retrieve identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.password, app.seed)
	if err != nil {
		return err
	}

	err = identity.Buckets().Add(identityBucket)
	if err != nil {
		return err
	}

	return app.identityService.Update(identity.Hash(), identity, app.password, app.password)
}

// Delete deletes a bucket from the given path
func (app *application) Delete(absolutePath string) error {
	// retrieve identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.password, app.seed)
	if err != nil {
		return err
	}

	buckets := identity.Buckets().All()
	for _, oneBucket := range buckets {
		if oneBucket.AbsolutePath() == absolutePath {
			err := identity.Buckets().Delete(oneBucket)
			if err != nil {
				return err
			}

			continue
		}
	}

	return app.identityService.Update(identity.Hash(), identity, app.password, app.password)
}

// Retrieve retrieves a bucket by hash
func (app *application) Retrieve(hash hash.Hash) (buckets.Bucket, error) {
	return app.bucketRepository.Retrieve(hash)
}

func (app *application) dirToFiles(rootPath string, relativePath string) ([]files.File, error) {
	path := filepath.Join(rootPath, relativePath)
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	out := []files.File{}
	for _, oneFile := range dirFiles {
		name := oneFile.Name()
		filePath := filepath.Join(relativePath, name)
		if oneFile.IsDir() {
			subFiles, err := app.dirToFiles(rootPath, filePath)
			if err != nil {
				return nil, err
			}

			out = append(out, subFiles...)
			continue
		}

		file, err := app.dirFileToFile(rootPath, filePath)
		if err != nil {
			return nil, err
		}

		out = append(out, file)
	}

	return out, nil
}

func (app *application) dirFileToFile(rootPath string, relativePath string) (files.File, error) {
	path := filepath.Join(rootPath, relativePath)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	index := 0
	chunks := []chunks.Chunk{}
	loops := int(math.Floor(float64(len(data)) / float64(app.chunkSizeInBytes)))
	for i := 0; i < loops; i++ {
		beginsOn := i * int(app.chunkSizeInBytes)
		createdOn := time.Now().UTC()
		if (i + 1) == loops {
			dataChk := data[beginsOn:]
			sizeInBytes := len(dataChk)
			dataHash, err := app.hashAdapter.FromBytes(dataChk)
			if err != nil {
				return nil, err
			}

			chk, err := app.chunkBuilder.Create().WithSizeInBytes(uint(sizeInBytes)).WithData(*dataHash).CreatedOn(createdOn).Now()
			if err != nil {
				return nil, err
			}

			chunks = append(chunks, chk)
			continue
		}

		stopsOn := (i + 1) * int(app.chunkSizeInBytes)
		dataChk := data[beginsOn:stopsOn]
		sizeInBytes := len(dataChk)
		dataHash, err := app.hashAdapter.FromBytes(dataChk)
		if err != nil {
			return nil, err
		}

		chk, err := app.chunkBuilder.Create().WithSizeInBytes(uint(sizeInBytes)).WithData(*dataHash).CreatedOn(createdOn).Now()
		if err != nil {
			return nil, err
		}

		chunks = append(chunks, chk)
		index++
	}

	createdOn := time.Now().UTC()
	return app.fileBuilder.Create().WithRelativePath(relativePath).WithChunks(chunks).CreatedOn(createdOn).Now()
}
