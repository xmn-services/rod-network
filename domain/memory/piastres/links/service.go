package links

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
	transfer_link "github.com/xmn-services/rod-network/domain/transfers/piastres/links"
)

type service struct {
	adapter      Adapter
	repository   Repository
	blockService blocks.Service
	trService    transfer_link.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	blockService blocks.Service,
	trService transfer_link.Service,
) Service {
	out := service{
		adapter:      adapter,
		repository:   repository,
		blockService: blockService,
		trService:    trService,
	}

	return &out
}

// Save saves a link
func (app *service) Save(link Link) error {
	hash := link.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	block := link.Next()
	err = app.blockService.Save(block)
	if err != nil {
		return err
	}

	trLink, err := app.adapter.ToTransfer(link)
	if err != nil {
		return err
	}

	return app.trService.Save(trLink)
}
