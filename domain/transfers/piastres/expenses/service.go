package expenses

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

// Save saves a expense instance
func (app *service) Save(expense Expense) error {
	js, err := app.adapter.ToJSON(expense)
	if err != nil {
		return err
	}

	return app.fileService.Save(expense.Hash().String(), js)
}
