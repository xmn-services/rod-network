package states

import (
	transfer_state "github.com/xmn-services/rod-network/domain/transfers/piastres/states"
)

type service struct {
	adapter    Adapter
	repository Repository
	trService  transfer_state.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	trService transfer_state.Service,
) Service {
	out := service{
		adapter:    adapter,
		repository: repository,
		trService:  trService,
	}

	return &out
}

// Save saves a state instance
func (app *service) Save(state State) error {
	chain := state.Chain()
	height := state.Height()
	retState, err := app.repository.Retrieve(chain, height)

	// if there is a state for the chain at the given height, but there is less
	// trx in the saved instance than the current one, update it:
	if err == nil {
		amountTrx := len(state.Transactions())
		retAmountTrx := len(retState.Transactions())
		if amountTrx <= retAmountTrx {
			return nil
		}
	}

	trState, err := app.adapter.ToTransfer(state)
	if err != nil {
		return err
	}

	return app.trService.Save(trState)
}
