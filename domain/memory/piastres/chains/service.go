package chains

import (
	"errors"
	"fmt"

	mined_block "github.com/xmn-services/rod-network/domain/memory/piastres/blocks/mined"
	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	mined_link "github.com/xmn-services/rod-network/domain/memory/piastres/links/mined"
	transfer_chains "github.com/xmn-services/rod-network/domain/transfers/piastres/chains"
)

type service struct {
	adapter        Adapter
	repository     Repository
	genesisService genesis.Service
	blockService   mined_block.Service
	linkService    mined_link.Service
	trService      transfer_chains.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	genesisService genesis.Service,
	blockService mined_block.Service,
	linkService mined_link.Service,
	trService transfer_chains.Service,
) Service {
	out := service{
		adapter:        adapter,
		repository:     repository,
		genesisService: genesisService,
		blockService:   blockService,
		linkService:    linkService,
		trService:      trService,
	}

	return &out
}

// Update updates a chain
func (app *service) Update(original Chain, updated Chain) error {
	// make the the genesis is the same in both chains:
	updatedGenHash := updated.Genesis().Hash()
	originalGenHash := original.Genesis().Hash()
	if originalGenHash.Compare(updatedGenHash) {
		str := fmt.Sprintf("the chain cannot be updated at height (%d) because its Genesis instance is invalid (updated: %s, stored: %s)", original.Head().Link().Index(), updatedGenHash.String(), originalGenHash.String())
		return errors.New(str)
	}

	// make sure the root is the same in both chains:
	updatedRootHash := updated.Root().Hash()
	originalRootHash := original.Root().Hash()
	if originalRootHash.Compare(updatedRootHash) {
		str := fmt.Sprintf("the chain cannot be updated at height (%d) because its Root mined Block instance is invalid (updated: %s, stored: %s)", original.Head().Link().Index(), updatedRootHash.String(), originalRootHash.String())
		return errors.New(str)
	}

	return app.save(updated)
}

// Insert inserts a chain
func (app *service) Insert(chain Chain) error {
	_, err := app.repository.Retrieve()
	if err == nil {
		return nil
	}

	// retrieve data:
	gen := chain.Genesis()
	root := chain.Root()

	// save genesis:
	err = app.genesisService.Save(gen)
	if err != nil {
		return err
	}

	// save root:
	err = app.blockService.Save(root)
	if err != nil {
		return err
	}

	return app.save(chain)
}

func (app *service) save(chain Chain) error {
	// retrieve data:
	head := chain.Head()

	// save head:
	err := app.linkService.Save(head)
	if err != nil {
		return err
	}

	// save the transfer chain:
	trChain, err := app.adapter.ToTransfer(chain)
	if err != nil {
		return err
	}

	return app.trService.Save(trChain)
}
