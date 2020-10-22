package genesis

import (
	"errors"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	transfer_genesis "github.com/xmn-services/rod-network/domain/transfers/piastres/genesis"
)

type service struct {
	adapter     Adapter
	repository  Repository
	billService bills.Service
	trService   transfer_genesis.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	billService bills.Service,
	trService transfer_genesis.Service,
) Service {
	out := service{
		adapter:     adapter,
		repository:  repository,
		billService: billService,
		trService:   trService,
	}

	return &out
}

// Save saves a genesis instance
func (app *service) Save(genesis Genesis) error {
	_, err := app.repository.Retrieve()
	if err == nil {
		return errors.New("there is already a Genesis instance")
	}

	bill := genesis.Bill()
	err = app.billService.Save(bill)
	if err != nil {
		return err
	}

	trGenesis, err := app.adapter.ToTransfer(genesis)
	if err != nil {
		return err
	}

	return app.trService.Save(trGenesis)
}
