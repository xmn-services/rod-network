package informations

import (
	"github.com/xmn-services/rod-network/libs/hashtree"
	transfer_information "github.com/xmn-services/rod-network/domain/transfers/buckets/informations"
)

type adapter struct {
	hashTreeBuilder hashtree.Builder
	trBuilder       transfer_information.Builder
}

func createAdapter(
	hashTreeBuilder hashtree.Builder,
	trBuilder transfer_information.Builder,
) Adapter {
	out := adapter{
		hashTreeBuilder: hashTreeBuilder,
		trBuilder:       trBuilder,
	}

	return &out
}

// ToTransfer converts an information to a transfer information
func (app *adapter) ToTransfer(information Information) (transfer_information.Information, error) {
	hash := information.Hash()
	files := information.Files()

	blocks := [][]byte{}
	for _, oneFile := range files {
		blocks = append(blocks, oneFile.Hash().Bytes())
	}

	ht, err := app.hashTreeBuilder.Create().WithBlocks(blocks).Now()
	if err != nil {
		return nil, err
	}

	amount := uint(len(files))
	createdOn := information.CreatedOn()
	builder := app.trBuilder.Create().WithHash(hash).WithFiles(ht).WithAmount(amount).CreatedOn(createdOn)
	if information.HasParent() {
		parent := information.Parent().Hash()
		builder.WithParent(parent)
	}

	return builder.Now()
}
