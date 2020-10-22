package files

import (
	"github.com/xmn-services/rod-network/libs/hashtree"
	transfer_file "github.com/xmn-services/rod-network/domain/transfers/buckets/files"
)

type adapter struct {
	hashTreeBuilder hashtree.Builder
	trBuilder       transfer_file.Builder
}

func createAdapter(
	hashTreeBuilder hashtree.Builder,
	trBuilder transfer_file.Builder,
) Adapter {
	out := adapter{
		hashTreeBuilder: hashTreeBuilder,
		trBuilder:       trBuilder,
	}

	return &out
}

// ToTransfer converts a file to a transfer file
func (app *adapter) ToTransfer(file File) (transfer_file.File, error) {
	hash := file.Hash()
	relativePath := file.RelativePath()
	chunks := file.Chunks()

	blocks := [][]byte{}
	for _, oneChunk := range chunks {
		blocks = append(blocks, oneChunk.Hash().Bytes())
	}

	ht, err := app.hashTreeBuilder.Create().WithBlocks(blocks).Now()
	if err != nil {
		return nil, err
	}

	amount := uint(len(chunks))
	createdOn := file.CreatedOn()
	return app.trBuilder.Create().
		WithHash(hash).
		WithRelativePath(relativePath).
		WithChunks(ht).
		WithAmount(amount).
		CreatedOn(createdOn).
		Now()
}
