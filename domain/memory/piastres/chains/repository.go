package chains

import (
	"errors"
	"fmt"

	mined_block "github.com/xmn-services/rod-network/domain/memory/piastres/blocks/mined"
	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	mined_link "github.com/xmn-services/rod-network/domain/memory/piastres/links/mined"
	transfer_chains "github.com/xmn-services/rod-network/domain/transfers/piastres/chains"
)

type repository struct {
	genesisRepository genesis.Repository
	blockRepository   mined_block.Repository
	linkRepository    mined_link.Repository
	trRepository      transfer_chains.Repository
	builder           Builder
}

func createRepository(
	genesisRepository genesis.Repository,
	blockRepository mined_block.Repository,
	linkRepository mined_link.Repository,
	trRepository transfer_chains.Repository,
	builder Builder,
) Repository {
	out := repository{
		genesisRepository: genesisRepository,
		blockRepository:   blockRepository,
		linkRepository:    linkRepository,
		trRepository:      trRepository,
		builder:           builder,
	}

	return &out
}

// Retrieve retrieves a chain instance
func (app *repository) Retrieve() (Chain, error) {
	trChain, err := app.trRepository.Retrieve()
	if err != nil {
		return nil, err
	}

	genHash := trChain.Genesis()
	gen, err := app.genesisRepository.Retrieve()
	if err != nil {
		return nil, err
	}

	if genHash.Compare(gen.Hash()) {
		str := fmt.Sprintf("the stored genesis hash, in the chain's stored data is invalid (expected: %s, returned: %s)", genHash.String(), gen.Hash().String())
		return nil, errors.New(str)
	}

	rootHash := trChain.Root()
	root, err := app.blockRepository.Retrieve(rootHash)
	if err != nil {
		return nil, err
	}

	headHash := trChain.Head()
	head, err := app.linkRepository.Retrieve(headHash)
	if err != nil {
		return nil, err
	}

	total := trChain.Total()
	createdOn := trChain.CreatedOn()
	return app.builder.Create().
		WithGenesis(gen).
		WithRoot(root).
		WithHead(head).
		WithTotal(total).
		CreatedOn(createdOn).
		Now()
}
