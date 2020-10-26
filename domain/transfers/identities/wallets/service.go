package wallets

import (
	"github.com/xmn-services/rod-network/libs/file"
)

type service struct {
	adapter     Adapter
	fileService file.Service
}

func createService(adapter Adapter, fileService file.Service) Service {
	out := service{
		adapter:     adapter,
		fileService: fileService,
	}

	return &out
}

// Save saves a wallet instance
func (app *service) Save(wallet Wallet) error {
	js, err := app.adapter.ToJSON(wallet)
	if err != nil {
		return err
	}

	return app.fileService.Save(wallet.Hash().String(), js)
}

// Delete deletes a wallet instance
func (app *service) Delete(wallet Wallet) error {
	return app.fileService.Delete(wallet.Hash().String())
}
