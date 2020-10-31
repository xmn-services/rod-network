package bills

import (
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	transfer_bill "github.com/xmn-services/rod-network/domain/transfers/piastres/bills"
)

type service struct {
	adapter     Adapter
	repository  Repository
	lockService locks.Service
	trService   transfer_bill.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	lockService locks.Service,
	trService transfer_bill.Service,
) Service {
	out := service{
		adapter:     adapter,
		repository:  repository,
		lockService: lockService,
		trService:   trService,
	}

	return &out
}

// Save saves a bill instance
func (app *service) Save(bill Bill) error {
	hash := bill.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	lock := bill.Lock()
	err = app.lockService.Save(lock)
	if err != nil {
		return err
	}

	trBill, err := app.adapter.ToTransfer(bill)
	if err != nil {
		return err
	}

	return app.trService.Save(trBill)
}

// SaveAll saves a list of bills
func (app *service) SaveAll(bills []Bill) error {
	for _, oneBill := range bills {
		err := app.Save(oneBill)
		if err != nil {
			return err
		}
	}

	return nil
}
