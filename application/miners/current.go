package miners

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
	mined_blocks "github.com/xmn-services/rod-network/domain/memory/piastres/blocks/mined"
	"github.com/xmn-services/rod-network/domain/memory/piastres/chains"
	"github.com/xmn-services/rod-network/domain/memory/piastres/links"
	mined_links "github.com/xmn-services/rod-network/domain/memory/piastres/links/mined"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/hash"
)

type current struct {
	hashAdapter         hash.Adapter
	trxRepository       transactions.Repository
	chainRepository     chains.Repository
	blockBuilder        blocks.Builder
	blockRepository     blocks.Repository
	minedBlockBuilder   mined_blocks.Builder
	minedBlockService   mined_blocks.Service
	linkBuilder         links.Builder
	minedLinkBuilder    mined_links.Builder
	minedLinkRepository mined_links.Repository
	minedLinkService    mined_links.Service
}

func createCurrent(
	hashAdapter hash.Adapter,
	trxRepository transactions.Repository,
	chainRepository chains.Repository,
	blockBuilder blocks.Builder,
	blockRepository blocks.Repository,
	minedBlockBuilder mined_blocks.Builder,
	minedBlockService mined_blocks.Service,
	linkBuilder links.Builder,
	minedLinkBuilder mined_links.Builder,
	minedLinkRepository mined_links.Repository,
	minedLinkService mined_links.Service,
) Current {
	out := current{
		hashAdapter:         hashAdapter,
		trxRepository:       trxRepository,
		chainRepository:     chainRepository,
		blockBuilder:        blockBuilder,
		blockRepository:     blockRepository,
		minedBlockBuilder:   minedBlockBuilder,
		minedBlockService:   minedBlockService,
		linkBuilder:         linkBuilder,
		minedLinkBuilder:    minedLinkBuilder,
		minedLinkRepository: minedLinkRepository,
		minedLinkService:    minedLinkService,
	}

	return &out
}

// Block mines a block
func (app *current) Block(address string, trx []string) error {
	addressHash, err := app.hashAdapter.FromString(address)
	if err != nil {
		return err
	}

	trxHashes := []hash.Hash{}
	for _, oneTrx := range trx {
		trxHash, err := app.hashAdapter.FromString(oneTrx)
		if err != nil {
			return err
		}

		trxHashes = append(trxHashes, *trxHash)
	}

	transactions, err := app.trxRepository.RetrieveAll(trxHashes)
	if err != nil {
		return err
	}

	chain, err := app.chainRepository.Retrieve()
	if err != nil {
		return err
	}

	// calculate the difficulty:
	difficulty := difficulty(chain, uint(len(transactions)))

	// build the block:
	createdOn := time.Now().UTC()
	gen := chain.Genesis()
	block, err := app.blockBuilder.Create().
		WithAddress(*addressHash).
		WithGenesis(gen).
		WithTransactions(transactions).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		return err
	}

	// mine the block:
	minedCreatedOn := time.Now().UTC()
	results, err := mine(app.hashAdapter, difficulty, block.Hash())
	if err != nil {
		return err
	}

	minedBlock, err := app.minedBlockBuilder.Create().
		WithBlock(block).
		WithMining(results).
		CreatedOn(minedCreatedOn).
		Now()

	if err != nil {
		return err
	}

	return app.minedBlockService.Save(minedBlock)
}

// Link mines a link
func (app *current) Link(prevMinedLink string, nextBlock string) error {
	chain, err := app.chainRepository.Retrieve()
	if err != nil {
		return err
	}

	prevMinedLinkHash, err := app.hashAdapter.FromString(prevMinedLink)
	if err != nil {
		return err
	}

	nextBlockHash, err := app.hashAdapter.FromString(nextBlock)
	if err != nil {
		return err
	}

	prevMinedLnk, err := app.minedLinkRepository.Retrieve(*prevMinedLinkHash)
	if err != nil {
		return err
	}

	nxtBlock, err := app.blockRepository.Retrieve(*nextBlockHash)
	if err != nil {
		return err
	}

	prev := prevMinedLnk.Hash()
	linkCreatedOn := time.Now().UTC()
	link, err := app.linkBuilder.Create().
		WithPreviousLink(prev).
		WithNext(nxtBlock).
		CreatedOn(linkCreatedOn).
		Now()

	if err != nil {
		return err
	}

	// mine:
	difficulty := chain.Genesis().Difficulty().Link()
	results, err := mine(app.hashAdapter, difficulty, link.Hash())
	if err != nil {
		return err
	}

	// return the mined link:
	minedLinkCreatedOn := time.Now().UTC()
	minedLink, err := app.minedLinkBuilder.Create().
		WithLink(link).
		WithMining(results).
		CreatedOn(minedLinkCreatedOn).
		Now()

	if err != nil {
		return err
	}

	//save the mined link:
	return app.minedLinkService.Save(minedLink)
}
