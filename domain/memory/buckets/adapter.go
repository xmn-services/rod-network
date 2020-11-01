package buckets

import (
	transfer_bucket "github.com/xmn-services/rod-network/domain/transfers/buckets"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

type adapter struct {
	hashTreeBuilder hashtree.Builder
	trBuilder       transfer_bucket.Builder
}

func createAdapter(
	hashTreeBuilder hashtree.Builder,
	trBuilder transfer_bucket.Builder,
) Adapter {
	out := adapter{
		hashTreeBuilder: hashTreeBuilder,
		trBuilder:       trBuilder,
	}

	return &out
}

// ToTransfer converts an bucket to a transfer bucket
func (app *adapter) ToTransfer(bucket Bucket) (transfer_bucket.Bucket, error) {
	hash := bucket.Hash()
	files := bucket.Files()

	blocks := [][]byte{}
	for _, oneFile := range files {
		blocks = append(blocks, oneFile.Hash().Bytes())
	}

	ht, err := app.hashTreeBuilder.Create().WithBlocks(blocks).Now()
	if err != nil {
		return nil, err
	}

	amount := uint(len(files))
	createdOn := bucket.CreatedOn()
	return app.trBuilder.Create().WithHash(hash).WithFiles(ht).WithAmount(amount).CreatedOn(createdOn).Now()
}
